import React, { useState, useEffect } from 'react';
import Tree from './tree';
import axios from 'axios';

const ProcessTree = () => {
    const baseUrl = "/api";
    // const baseUrl = "http://localhost:8080";
    const baseUrlCpu = "cpu";

    const [cpu, setCpu] = useState({});
    const [process, setProcess] = useState({})


    useEffect(() => {
      const getCpu = () => {
        axios.get(`${baseUrl}/${baseUrlCpu}/get`)
          .then((response) => {
            setCpu(response.data);
          })
          .catch(error => {
            console.error("Error al obtener los datos del cpu", error);
          });
      }
      getCpu();
      console.log("A")
    });

    const handleSelectChange = (event) => {
      setProcess(event.target.value);
    }

  return (
    <div className="content">
      <h2>Arbol de Processos</h2>
      <h3>Selecciona una opcion</h3>
      <select value={process} onChange={handleSelectChange}>
        {cpu && cpu.processes.map((proc, index) => {
          return (
            <option key={index} value={proc}>{proc.pid}</option>
          )
        })}
      </select>
      <Tree process={process}/>
    </div>
  );
};

export default ProcessTree;
