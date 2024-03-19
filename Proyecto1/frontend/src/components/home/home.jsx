import React from 'react';
import { Link } from 'react-router-dom';
import './home.css'; // Importa tu archivo CSS para estilos

function Home() {
  return (
    <div className="navbar">
      <Link to="/real-time" className="nav-item">Real Time</Link>
      <Link to="/historical" className="nav-item">Historical</Link>
      <Link to="/processTree" className="nav-item">Process Tree</Link>
      <Link to="/processStateSimulation" className="nav-item">Process State Simulation</Link>
    </div>
  );
}

export default Home;
