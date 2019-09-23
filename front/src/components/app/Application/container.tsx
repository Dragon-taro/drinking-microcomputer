import { RootState } from "../../../entity/state";
import { StateProps, DispatchProps } from "./type";
import { connect } from "react-redux";
import Application from "./presentation";
import { getParty } from "../../../redux/party/effects";

const mapStateToProps = (state: RootState): StateProps => ({
  party: state.party
});

const mapDispatchToProps = (dispatch: any): DispatchProps => ({
  getParty: () => {
    dispatch(getParty());
  }
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Application);
