import * as React from "react";
import { Props } from "./type";

const Index: React.FunctionComponent<Props> = ({ finishParty, getData }) => {
  React.useEffect(() => {
    getData();
  }, []);

  return (
    <div>
      <div>
        <button onClick={finishParty}>飲み会を終了する</button>
      </div>
    </div>
  );
};

export default Index;
