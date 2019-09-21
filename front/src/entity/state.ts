import { Party } from "./party";
import { Data } from "./data";

export interface RootState {
  party: Party;
  data: Data[];
}
