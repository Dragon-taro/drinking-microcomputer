import { RootState } from "../../../entity/state";
import { StateProps, DispatchProps } from "./type";
import { connect } from "react-redux";
import Application from "./presentation";
import { getParty } from "../../../redux/party/effects";
import { getData, subscribeData } from "../../../redux/data/effects";

const mapStateToProps = (state: RootState): StateProps => ({
  party: state.party,
  data: state.data
});

const mapDispatchToProps = (dispatch: any): DispatchProps => ({
  getParty: () => {
    dispatch(getParty());
  },
  getData: () => {
    dispatch(getData());
  },
  subscribeData: () => {
    dispatch(subscribeData());
  }
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Application);
