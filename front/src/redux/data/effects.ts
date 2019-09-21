import { Dispatch } from "redux";
import { get } from "../utils/api";
import { setData } from "./action";
import { Data } from "../../entity/data";

export const getData = () => async (dispatch: Dispatch) => {
  const data = await get<Data[]>("data");

  dispatch(setData(data));
};
