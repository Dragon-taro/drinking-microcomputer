import * as React from "react";
import { Props } from "./type";
import { Line, ChartData } from "react-chartjs-2";

const Index: React.FunctionComponent<Props> = ({
  finishParty,
  data: _data,
  labels
}) => {
  const datasets: Chart.ChartDataSets[] = [
    {
      label: "飲んだ量",
      fill: false,
      lineTension: 0,
      backgroundColor: "rgba(75,192,192,0.4)",
      borderColor: "rgba(75,192,192,1)",
      borderCapStyle: "butt",
      borderDash: [],
      borderDashOffset: 0.0,
      borderJoinStyle: "miter",
      pointBorderColor: "rgba(75,192,192,1)",
      pointBackgroundColor: "#fff",
      pointBorderWidth: 1,
      pointHoverBackgroundColor: "rgba(75,192,192,1)",
      pointHoverBorderColor: "rgba(220,220,220,1)",
      pointHoverBorderWidth: 2,
      pointRadius: 1,
      data: _data
    }
  ];

  const options: Chart.ChartOptions = {
    animation: {
      easing: "easeOutQuart"
    },
    scales: {
      xAxes: [
        {
          type: "time",
          distribution: "linear",
          time: {
            displayFormats: {
              minute: "hh:mm"
            }
          }
        }
      ]
    }
  };

  const data: ChartData<Chart.ChartData> = {
    labels,
    datasets
  };

  return (
    <div>
      <Line data={data} options={options} />
      <div>
        <button onClick={finishParty}>飲み会を終了する</button>
      </div>
    </div>
  );
};

export default Index;
