import { Dispatch } from "redux";
import { get } from "../utils/api";
import { Party } from "../../entity/party";
import { setParty } from "./action";

export const getParty = () => async (dispatch: Dispatch) => {
  const party = await get<Party>("party");

  dispatch(setParty(party));
};
