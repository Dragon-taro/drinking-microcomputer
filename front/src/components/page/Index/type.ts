import { DataList } from "../../../entity/data";

export interface StateProps {
  data: DataList;
}

export interface DispatchProps {
  finishParty: () => void;
}

export interface MergedProps {
  labels: string[];
  data: number[];
}

export type Props = MergedProps & DispatchProps;
