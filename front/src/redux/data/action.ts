import { SET_DATA } from "./constant";
import { Action } from "../utils/action";
import { Data } from "../../entity/data";

export const setData = (payload: Data[] | null): Action<Data[] | null> => ({
  type: SET_DATA,
  payload: payload
});