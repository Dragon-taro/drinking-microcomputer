import { Party } from "../../../entity/party";
import { DataList } from "../../../entity/data";

export interface StateProps {
  party: Party;
  data: DataList;
}

export interface DispatchProps {
  getParty: () => void;
  getData: () => void;
}

export type Props = StateProps & DispatchProps;
