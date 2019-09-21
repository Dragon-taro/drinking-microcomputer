export interface StateProps {}

export interface DispatchProps {
  postParty: () => void;
}

export type Props = StateProps & DispatchProps;
