import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import LineGraph from './home/lineGraph';
import axios from 'axios';
import "./.css"


const HistoricalMonitoring = () => {
    const baseUrl = "/api";
    // const baseUrl = "http://localhost:8080";
    const baseUrlRam = "ram";
    const baseUrlCpu = "cpu";

    const [ram, setRam] = useState([]);
    const [ramTime, setRamTime] = useState([]);

    const GetAllRam = () => {
        axios.get(`${baseUrl}/${baseUrlRam}/getAll`)
            .then(response => {
                const data = response.data;
                const ramUsage = data.map((ram) => {
                  return ram.usedMemory;
                })
                setRam(ramUsage.reverse());
                const timeArray = data.map((e) => {return new Date(e.fecha).toLocaleDateString()});
                setRamTime(timeArray.reverse())

                // console.log(ram);
            })
            .catch(error => {
                console.error("Error al obtener los datos de la ram", error);
            });
    }

    const [cpu, setCpu] = useState([])
    const [cpuTime, setCpuTime] = useState([]);

    const GetAllCpu = () => {
        axios.get(`${baseUrl}/${baseUrlCpu}/getAll`)
            .then(response => {
                const data = response.data;
                const cpuUsage = data.map((cpu) => {
                  return cpu.cpu_percentage;
                })
                setCpu(cpuUsage.reverse())
                const timeArray = data.map((e) => {return new Date(e.fecha).toLocaleDateString()});
                setRamTime(timeArray.reverse())
                setCpuTime(timeArray.reverse())

                // console.log(cpu);
            })
            .catch(error => {
                console.error("Error al obtener los datos de la ram", error);
            });
    }

    useEffect(() => {
        const GetRamInterval = setInterval(GetAllCpu, 500);
        const GetCpuInterval = setInterval(GetAllRam, 500);

        return () => {
            clearInterval(GetRamInterval);
            clearInterval(GetCpuInterval);
        }
    }, []);

  return (
     <div className="app-container">
      <div className="centered-container">
        <h1>Monitoreo Historico</h1>
        <div className='graph-container'>
            <LineGraph titulo="RAM" label="Uso de la memoria RAM a lo largo del tiempo" tiempos={ramTime} valores={ram}/>
            <LineGraph titulo="CPU" label="Uso del CPU a lo largo del tiempo" tiempos={cpuTime} valores={cpu}/>
        </div>
        <Link to="/">Regresar a la pagina principal</Link>
      </div>
    </div>
  );
};

export default HistoricalMonitoring;
