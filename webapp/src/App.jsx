import './index.css'
import React from 'react';
import { createRoot } from 'react-dom/client';

function App() {
  return (
    <div className="mr-2">
      <h1 className="text-3xl font-extrabold font-sans selection:bg-yellow-400">How will your climate change?</h1>
    </div>
  );
}

const container = document.getElementById('root');
const root = createRoot(container); // createRoot(container!) if you use TypeScript
root.render(<App/>);

export default App;