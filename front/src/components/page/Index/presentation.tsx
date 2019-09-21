import * as React from "react";
import { Props } from "./type";

const Index: React.FunctionComponent<Props> = ({
  finishParty,
  getData,
  data
}) => {
  React.useEffect(() => {
    getData();
  }, []);

  return (
    <div>
      <div>
        {data &&
          data.map(d => {
            return (
              <div>
                <div>作成日時:{d.createdAt}</div>
                <div>量: {d.amount}</div>
                <div>合計量: {d.totalAmount}</div>
                <div>前回との差分: {d.diff}</div>
                <hr />
              </div>
            );
          })}
      </div>
      <div>
        <button onClick={finishParty}>飲み会を終了する</button>
      </div>
    </div>
  );
};

export default Index;
