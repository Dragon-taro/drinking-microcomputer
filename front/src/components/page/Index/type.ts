export interface StateProps {}

export interface DispatchProps {
  finishParty: () => void;
}

export type Props = StateProps & DispatchProps;
