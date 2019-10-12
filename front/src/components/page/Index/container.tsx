import { connect } from "react-redux";

import { finishParty } from "../../../redux/party/effects";
import NewParty from "./presentation";
import { DispatchProps, Props } from "./type";
import { getData } from "../../../redux/data/effects";
import { RootState } from "../../../entity/state";
import { StateProps } from "./type";
import { Data } from "../../../entity/data";

const mapStateToProps = (state: RootState): StateProps => ({
  data: state.data
});

const mapDispatchToProps = (dispatch: any): DispatchProps => ({
  finishParty: () => {
    dispatch(finishParty());
  },
  getData: () => {
    dispatch(getData());
  }
});

const mergeProps = (state: StateProps, dispatch: DispatchProps): Props => {
  const { data: _data } = state;

  const labels = filterData(d => d.createdAt, _data);
  const data = filterData(d => d.totalAmount, _data);

  return {
    labels,
    data,
    ...dispatch
  };
};

const filterData = <T,>(f: (d: Data) => T, data: Data[] | null): T[] => {
  return data ? data.map(f).reverse() : [];
};

export default connect(
  mapStateToProps,
  mapDispatchToProps,
  mergeProps
)(NewParty);
