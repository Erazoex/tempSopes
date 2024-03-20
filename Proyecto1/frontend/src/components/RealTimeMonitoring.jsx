import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import PieGraph from './pieGraph';
import axios from 'axios';
import "./.css"


const RealTimeMonitoring = () => {
    const baseUrl = "/api";
    // const baseUrl = "http://localhost:8080";
    const baseUrlRam = "ram";
    const baseUrlCpu = "cpu";

    const [ram, setRam] = useState({});

    const GetRam = () => {
        axios.get(`${baseUrl}/${baseUrlRam}/get`)
            .then(response => {
                setRam(response.data);
                // console.log(ram);
            })
            .catch(error => {
                console.error("Error al obtener los datos de la ram", error);
            });
    }

    const [cpu, setCpu] = useState({})

    const GetCpu = () => {
        axios.get(`${baseUrl}/${baseUrlCpu}/get`)
            .then(response => {
                setCpu(response.data);
                // console.log(cpu);
            })
            .catch(error => {
                console.error("Error al obtener los datos del cpu", error);
            });
    }

    useEffect(() => {
        const GetRamInterval = setInterval(GetRam, 500);
        const GetCpuInterval = setInterval(GetCpu, 500);

        return () => {
            clearInterval(GetRamInterval);
            clearInterval(GetCpuInterval);
        }
    }, []);

  return (
     <div className="app-container">
      <div className="centered-container">
        <h1>Monitoreo en Tiempo Real</h1>
        <div className='graph-container'>
            <PieGraph titulo="RAM" label="Uso de la memoria RAM" EnUso={ram?.usedMemory} Total={ram?.totalRam}/>
            <PieGraph titulo="CPU" label="Uso del CPU" EnUso={cpu?.cpu_percentage} Total={cpu?.cpu_total}/>
        </div>
        <Link to="/">
          <button>Regresar a la pagina principal</button>
        </Link>
      </div>
    </div>
  );
};

export default RealTimeMonitoring;
