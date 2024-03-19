import React from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Doughnut } from "react-chartjs-2";

ChartJS.register(ArcElement, Tooltip, Legend);

const PieGraph = ({ titulo, label , EnUso, Total}) => {
  const data = {
    labels: ['Espacio Libre', 'Espacio en Uso'],
    datasets: [
      {
        label: label,
        data: [Total - EnUso, EnUso],
        backgroundColor: ['rgb(54, 162, 235)', 'rgb(255, 99, 132)'],
        hoverOffset: 4,
      }],
  };

  return (
    <div>
      <h2>{titulo}</h2>
      <p>{((EnUso/ Total) * 100).toFixed(2)}% en uso</p>
      <Doughnut data={data} />
    </div>
  );
};

export default PieGraph;