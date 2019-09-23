import { Data } from "../../../entity/data";

export interface StateProps {
  data: Data[] | null;
}

export interface DispatchProps {
  finishParty: () => void;
  getData: () => void;
}

export type Props = StateProps & DispatchProps;
