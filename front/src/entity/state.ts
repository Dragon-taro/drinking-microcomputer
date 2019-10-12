import { Party } from "./party";
import { DataList } from "./data";

export interface RootState {
  party: Party;
  data: DataList;
}
