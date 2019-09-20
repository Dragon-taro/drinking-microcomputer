import { RootState } from "../../../entity/state";
import { StateProps } from "./type";
import { connect } from "react-redux";
import Application from "./presentation";
import { Dispatch } from "redux";
import { getParty } from "../../../redux/party/effects";

const mapStateToProps = (state: RootState): StateProps => ({
  party: state.party
});

const mapDispatchToProps = (dispatch: any) => ({
  getParty: () => {
    dispatch(getParty());
  }
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Application);
