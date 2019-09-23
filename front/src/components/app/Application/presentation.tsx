import * as React from "react";
import { Props } from "./type";
import NewParty from "../../page/NewParty/container";
import Index from "../../page/Index/container";

const Application: React.FunctionComponent<Props> = ({ getParty, party }) => {
  React.useEffect(() => {
    getParty();
  }, []);

  if (party) {
    return <Index />;
  } else {
    return <NewParty />;
  }
};

export default Application;
