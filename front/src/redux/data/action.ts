import { SET_DATA, LODING_DATA, SET_SINGLE_DATA } from "./constant";
import { Action } from "../utils/action";
import { Data } from "../../entity/data";

export const setData = (payload: Data[] | null): Action<Data[] | null> => ({
  type: SET_DATA,
  payload: payload
});

export const setSingleData = (payload: Data): Action<Data> => ({
  type: SET_SINGLE_DATA,
  payload: payload
});

export const loadingData = (): Action<{}> => ({
  type: LODING_DATA
});
