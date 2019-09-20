import * as React from "react";
import { Props } from "./type";

const Application: React.FunctionComponent<Props> = ({ party, getParty }) => {
  React.useEffect(() => {
    getParty();
  }, []);
  console.log(party);

  return <div>hello</div>;
};

export default Application;
