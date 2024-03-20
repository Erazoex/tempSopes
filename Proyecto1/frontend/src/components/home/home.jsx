import React from 'react';
import { Link } from 'react-router-dom';
import './home.css'; // Importa tu archivo CSS para estilos

function Home() {
  return (
    <div className='App'>
      <div className="data-container">
        <h1>Universidad de San Carlos de Guatemala</h1>
        <h2>Laboratorio de Sistemas Operativos 1</h2>
        <h2>Brian Josue Erazo Sagastume - 201807253</h2>
      </div>
      <div className="button-container">
        <Link to="/real-time" className="nav-item"><button>Monitoreo en tiempo real</button></Link>
        <Link to="/historical" className="nav-item"><button>Monitoreo historico</button></Link>
        <Link to="/processTree" className="nav-item"><button>Arbol de procesos</button></Link>
        <Link to="/processStateSimulation" className="nav-item"><button>Diagrama de estados</button></Link>
      </div>

    </div>
  );
}

export default Home;
