import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement} from "chart.js";
import { Line } from "react-chartjs-2";
import '../App.css'

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement);


const LineGraph = ({ titulo, label , tiempos, valores}) => {
  const data = {
    labels: tiempos,
    datasets: [
      {
        label: label,
        data: valores,
        backgroundColor: 'rgb(54, 162, 235)',
        hoverOffset: 4,
      }],
  };

  const options = {
    scales: {
      x: {
        title: {
          display: true,
          text: 'Fechas'
        }
      },
      y: {
        title: {
          display: true,
          text: 'Uso'
        }
      }
    }
  };

  return (
    <div className="linechart-container">
      <h2>{titulo}</h2>
      <Line data={data} options={options} />
    </div>
  );
};

export default LineGraph;
