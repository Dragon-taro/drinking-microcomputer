import { Party } from "../../../entity/party";

export interface StateProps {
  party: Party;
}

interface DispatchProps {
  getParty: () => void;
}

export type Props = StateProps & DispatchProps;
