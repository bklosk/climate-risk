import React from 'react';
import { useEffect, useState } from 'react';
import {Stack} from '@visx/shape';
import { scaleLinear, scaleOrdinal } from '@visx/scale';
import { transpose } from '@visx/vendor/d3-array';
import { animated, useSpring } from '@react-spring/web';
import generateData from './generateData';
import useForceUpdate from './forceUpdate';

// utility function to generate a range of numbers
const range = (n: number) => Array.from(new Array(n), (_, i) => i);
const keys = range(20);
const samples = 100;
function randomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

// scales x and y of the stream appropriately
const xScale = scaleLinear<number>({
  domain: [0, samples - 1],
});
const yScale = scaleLinear<number>({
  domain: [-30, 50],
});

const layers = transpose<number>(
  keys.map(() => generateData(samples, randomInt(4, 20)))
)
// get y0 and y1 values from generated data
type Datum = number[];
const getY0 = (d: Datum) => yScale(d[0]) ?? 0;
const getY1 = (d: Datum) => yScale(d[1]) ?? 0;

// accessors for the streamgraph props
export type StreamGraphProps = {
  width: number;
  height: number;
  animate?: boolean;
};

const colorScale = scaleOrdinal<number, string>({
  domain: keys,
  range: ['#A3BAC3', '#EB5E28', '#363537', '#C8AD55', '#9ecadd', '#ABC0C8'],
});

export default function Streamgraph({width, height, animate = true}: StreamGraphProps) {  
  xScale.range([0, width]);
  yScale.range([height, 0]);
  const [time, setTime] = useState(Date.now());
  const forceUpdate = useForceUpdate();
  const force = () => forceUpdate();

  useEffect(() => {
    const interval = setInterval(() => {
      setTime(Date.now());
      force(); 
    }, 1);
    return () => {
      clearInterval(interval);
    };
  }, [force]);

  return (
    <svg width={width} height={height}>
      <g onClick={force}>
        <rect x={0} y={0} width={width} height={height} fill={'#FFFFF0'} rx={14} />
        <Stack<number[], number>
          data={layers}
          keys={keys}
          offset="wiggle"
          color={colorScale}
          x={(_, i) => xScale(i) ?? 0}
          y0={getY0}
          y1={getY1}
        >
          {({ stacks, path }) =>
            stacks.map((stack) => {
              // I have no idea how this works lmao
              // Alternatively use renderprops <Spring to={{ d }}>{tweened => ...}</Spring>
              const pathString = path(stack) || '';
              const tweened = animate ? useSpring({ pathString }) : { pathString };
              const color = colorScale(stack.key);
              return (
                <g key={`series-${stack.key}`}>
                  <animated.path d={tweened.pathString} fill={color} />
                  <animated.path d={tweened.pathString} fill={color} />
                </g>
              );
            })
          }
        </Stack>
      </g>

  </svg>
  )
};