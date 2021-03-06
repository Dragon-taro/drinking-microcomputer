import { Dispatch } from "redux";
import { get, post, patch } from "../utils/api";
import { Party } from "../../entity/party";
import { setParty, loadingParty } from "./action";

export const getParty = () => async (dispatch: Dispatch) => {
  dispatch(loadingParty());

  const party = await get<Party>("party");
  dispatch(setParty(party));
};

export const postParty = () => async (dispatch: Dispatch) => {
  dispatch(loadingParty());

  const party = await post<Party>("party");
  dispatch(setParty(party));
};

export const finishParty = () => async (dispatch: Dispatch) => {
  await patch<Party>("party/finish");

  dispatch(setParty(null));
};
