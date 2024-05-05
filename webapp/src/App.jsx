import './index.css'
import Streamgraph from "./components/Streamgraph";
import React from 'react';
import {useState} from 'react';
import { createRoot } from 'react-dom/client';

function App() {
  const [width, setWidth] = useState(window.innerWidth);
  const [height, setHeight] = useState(window.innerHeight);
  function resize() {
    setWidth(window.innerWidth);
    setHeight(window.innerHeight);
  };
  window.addEventListener('resize', resize);

  return (
    <div className="">
      <h1 className="text-3xl font-extrabold font-sans flex justify-center pt-12 selection:bg-yellow-400">How will your climate change?</h1>
      <Streamgraph width={width} height={height / 3} className="align-bottom"/>
    </div>

  );
}

const container = document.getElementById('root');
const root = createRoot(container);
root.render(<App/>);