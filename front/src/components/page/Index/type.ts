import { Data } from "../../../entity/data";

export interface StateProps {
  data: Data[] | null;
}

export interface DispatchProps {
  finishParty: () => void;
  getData: () => void;
}

export interface MergedProps {
  labels: string[];
  data: number[];
}

export type Props = MergedProps & DispatchProps;
