import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import Tree from './tree';
import axios from 'axios';

const ProcessTree = () => {
    const baseUrl = "/api";
    // const baseUrl = "http://localhost:8080";
    const baseUrlCpu = "cpu";

    const [cpu, setCpu] = useState(null);
    const [process, setProcess] = useState(null);

    useEffect(() => {
      const getCpu =  async () => {
        try {
          const response = await axios.get(`${baseUrl}/${baseUrlCpu}/get`);
          setCpu(response.data);

        } catch (error) {
          console.error("Error al obtener los datos del cpu", error)
        }
      }
      getCpu();
    }, []);

    const handleSelectChange = (event) => {
      if (event.target.value == -1) {
        return;
      }
      const selectedPid = event.target.value;
      const selected = cpu.processes.find(process => process.pid == selectedPid);
      setProcess(selected);
    }

  return (
    <div className="content">
      <h1>Arbol de Processos</h1>
      {cpu && (
        <div>
          <select onChange={handleSelectChange}>
            <option value={-1}>Seleccione una opcion</option>
            {cpu.processes.map((process, index) => (
              <option key={index} value={process.pid}>
                {process.pid}
              </option>
            ))}
          </select>
          {process && <Tree process={process}/>}
        </div>
      )}
      <Link to="/">
        <button>Regresar a la pagina principal</button>
      </Link>
    </div>
  );
};

export default ProcessTree;
