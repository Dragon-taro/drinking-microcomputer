import * as React from "react";
import { Props } from "./type";
import NewParty from "../../page/NewParty/container";
import Index from "../../page/Index/container";
import Loading from "../../page/Loading";

const Application: React.FunctionComponent<Props> = ({
  getParty,
  getData,
  subscribeData,
  party,
  data
}) => {
  React.useEffect(() => {
    getParty();
    getData();
    subscribeData();
  }, []);

  const loading = party.isLoading || data.isLoading;

  if (loading) {
    return <Loading />;
  } else if (party.createdAt) {
    return <Index />;
  } else {
    return <NewParty />;
  }
};

export default Application;
