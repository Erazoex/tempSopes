import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Network } from 'vis-network/standalone';
import 'vis-network/styles/vis-network.css';
import axios from 'axios';
import "../App.css"

const ProcessStateSimulation = () => {
  const [processes, setProcesses] = useState(null);
  const [pid, setPid] = useState(null);
  const [status, setStatus] = useState('');
  const [network, setNetwork] = useState(null);

  const baseUrl = "/api";
  // const baseUrl = "http://localhost:8080";
  const baseUrlProcess = "process"

  const getProcesses =  async () => {
    try {
      const response = await axios.get(`${baseUrl}/${baseUrlProcess}/getAll`);
      setProcesses(response.data);

    } catch (error) {
      console.error("Error al obtener los datos del cpu", error)
    }
  }

  useEffect(() => {
    const container = document.getElementById('network');
    const options = {
      nodes: { borderWidth: 2 },
      edges: { width: 2 }
    };
    
    const nodes = [
      { id: 1, label: 'New', color: 'lightblue', originalColor: 'lightblue' },
      { id: 2, label: 'Ready', color: 'lightgreen', originalColor: 'lightgreen' },
      { id: 3, label: 'Waiting', color: 'lightgray', originalColor: 'lightgray' },
      { id: 4, label: 'Running', color: 'lightsalmon', originalColor: 'lightsalmon' },
      { id: 5, label: 'Terminated', color: 'lavender', originalColor: 'lavender' }
    ];

    const edges = [
      { from: 1, to: 2 },
      { from: 2, to: 4 },
      { from: 4, to: 2 },
      { from: 4, to: 3 },
      { from: 3, to: 2 },
      { from: 4, to: 5 }
    ];

    const newNetwork = new Network(container, { nodes, edges }, options);
    setNetwork(newNetwork);

    getProcesses();

    return () => {
      newNetwork.destroy();
    };
  }, []);

  const updateNetwork = (state) => {
    const nodes = network.body.data.nodes.get();
    const updatedNodes = nodes.map(node => {
      if (node.label === state) {
        return { ...node, color: 'green' }; 
      }
      return { ...node, color: node.originalColor }; 
    });
    network.body.data.nodes.update(updatedNodes);
  };

  const clearNetwork = () => {
    const nodes = network.body.data.nodes.get();
    const updatedNodes = nodes.map(node => {
      return { ...node, color: node.originalColor }; 
    });
    network.body.data.nodes.update(updatedNodes);
  }

  const newProcess = async () => {
    try {
      const response = await axios.get(`${baseUrl}/${baseUrlProcess}/new`);
      const data = response.data;
      setPid(data.pid);
      setStatus(data.state);
      updateNetwork(data.state);
      getProcesses();
      // console.log(data.state);
    } catch (error) {
      setStatus('Error al crear un nuevo proceso');
    }
  };

  const readyProcess = async () => {
    if (!pid) {
      setStatus('No existe un proceso seleccionado');
      return;
    }

    try {
      const response = await axios.get(`${baseUrl}/${baseUrlProcess}/ready?pid=${pid}`);
      const data = response.data;
      setPid(data.pid)
      setStatus(data.state);
      updateNetwork(data.state);
      getProcesses();
      // console.log(data.state);
    } catch (error) {
      setStatus('Error al cambiar el status a Ready');
    }
  };

  const waitingProcess = async () => {
    if (!pid) {
      setStatus('No existe un proceso seleccionado');
      return;
    }

    try {
      const response = await axios.get(`${baseUrl}/${baseUrlProcess}/waiting?pid=${pid}`);
      const data = response.data;
      setPid(data.pid)
      setStatus(data.state);
      updateNetwork(data.state);
      getProcesses();
      // console.log(data.state);
    } catch (error) {
      setStatus('Error al cambiar el status a Waiting');
    }
  };

  const runningProcess = async () => {
    if (!pid) {
      setStatus('No existe un proceso seleccionado');
      return;
    }

    try {
      const response = await axios.get(`${baseUrl}/${baseUrlProcess}/running?pid=${pid}`);
      const data = response.data;
      setPid(data.pid)
      setStatus(data.state);
      updateNetwork(data.state);
      getProcesses();
      // console.log(data.state);
    } catch (error) {
      setStatus('Error al cambiar el status a Running');
    }
  };

  const terminatedProcess = async () => {
    if (!pid) {
      setStatus('No existe un proceso seleccionado');
      return;
    }

    try {
      const response = await axios.get(`${baseUrl}/${baseUrlProcess}/terminated?pid=${pid}`);
      const data = response.data;
      setPid(data.pid)
      setStatus(data.state);
      updateNetwork(data.state);
      getProcesses();
      // console.log(data.state);
    } catch (error) {
      setStatus('Error al cambiar el status a Terminated');
    }
  };

  const handleSelectChange = (event) => {
    if (event.target.value == -1) {
      setStatus("Se debe escoger un proceso");
      clearNetwork();
      return;
    }
    const selectedPid = event.target.value;
    const selected = processes.find(process => process.pid == selectedPid);
    setPid(selected.pid);
    setStatus(selected.state);
    updateNetwork(selected.state);
    // console.log(data.state);
  }

  return (
    <div>
      <div className='process-container'>
        <h1 >Diagrama de Estados</h1>
      </div>
      <div className='process-container'>
        <button onClick={newProcess}>New</button>
        <button onClick={readyProcess}>Ready</button>
        <button onClick={waitingProcess}>Waiting</button>
        <button onClick={runningProcess}>Running</button>
        <button onClick={terminatedProcess}>Terminated</button>
      </div>
      <div className='process-container'>
        {processes && (
          <div>
            <select onChange={handleSelectChange}>
              <option value={-1}>Seleccione una opcion</option>
              {processes.map((process, index) => (
                <option key={index} value={process.pid}>
                  {process.pid}
                </option>
              ))}
            </select>
            <h3>Current State: {status}</h3>
          </div>
        )}
      </div>
      <div id="network" style={{ width: '600px', height: '400px' }}></div>
      <div className='process-container'>
        <Link to="/">
          <button>Regresar a la pagina principal</button>
        </Link>
      </div>
    </div>
  );
};

export default ProcessStateSimulation;
