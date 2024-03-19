import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import RealTimeMonitoring from './components/RealTimeMonitoring';
import HistoricalMonitoring from './components/HistoricalMonitoring';
import ProcessTree from './components/ProcessTree';
import ProcessStateSimulation from './components/ProcessStateSimulation';
import Home from './components/home/home';

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/real-time" element={<RealTimeMonitoring />}/>
        <Route path="/historical" element={<HistoricalMonitoring/>}/>
        <Route path="/processTree" element={<ProcessTree/>}/>
        <Route path="/processStateSimulation" element={<ProcessStateSimulation/>}/>
      </Routes>
    </BrowserRouter>
  );
};


export default App;