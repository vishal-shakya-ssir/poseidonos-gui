/* -------------------------------------------------------------------------------------/
                                                                                    /
/               COPYRIGHT (c) 2019 SAMSUNG ELECTRONICS CO., LTD.                      /
/                          ALL RIGHTS RESERVED                                        /
/                                                                                     /
/   Permission is hereby granted to licensees of Samsung Electronics Co., Ltd.        /
/   products to use or abstract this computer program for the sole purpose of         /
/   implementing a product based on Samsung Electronics Co., Ltd. products.           /
/   No other rights to reproduce, use, or disseminate this computer program,          /
/   whether in part or in whole, are granted.                                         / 
/                                                                                     /
/   Samsung Electronics Co., Ltd. makes no representation or warranties with          /
/   respect to the performance of this computer program, and specifically disclaims   /
/   any responsibility for any damages, special or consequential, connected           /
/   with the use of this program.                                                     /
/                                                                                     /
/-------------------------------------------------------------------------------------/


DESCRIPTION: <File description> *
@NAME : index.js
@AUTHORS: Jay Hitesh Sanghavi 
@Version : 1.0 *
@REVISION HISTORY
[03/06/2019] [Jay] : Prototyping..........////////////////////
*/
import React, { Component } from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import ThemeProvider from '@material-ui/core/styles/MuiThemeProvider';
import {
  Button,
  Paper,
  Grid,
  Typography,
  withStyles,
  TextField,
  FormControl,
  Select,
  MenuItem,
  Checkbox,
  FormControlLabel,
  Tooltip
} from '@material-ui/core';
import { customTheme, PageTheme } from '../../../theme';
import '../../../containers/Volume/Volume.css';
import './CreateVolume.css';
import AlertDialog from '../../Dialog';
import * as actionCreators from "../../../store/actions/exportActionCreators";
import * as actionTypes from "../../../store/actions/actionTypes";

const io = require('socket.io-client');

const styles = (theme) => ({
  formContainer: {
    width: '100%',
    display: 'flex',
    padding: theme.spacing(0, 4),
    flexWrap: 'wrap',
    boxSizing: 'border-box'
  },
  volBtnContainer: {
    margin: theme.spacing(1, 0)
  },
  unitSelect: {
    marginTop: theme.spacing(2),
    height: 32
  },
  unitText: {
    width: 'calc(80% - 60px)',
    display: 'flex',
    justifyContent: 'flex-end',
    [theme.breakpoints.down('xs')]: {
      width: '60%'
    }
  },
  formControl: {
    [theme.breakpoints.down('xs')]: {
      justifyContent: 'center'
    }
  },
  button: {
    height: '1.8rem',
    lineHeight: '0px'
  },
  volumeName: {
    width: '80%'
  },
  volumeUnit: {
    minWidth: 60,
    [theme.breakpoints.down('xs')]: {
      width: '20%'
    }
  },
  volumeCreatePaper: {
    height: 330,
    [theme.breakpoints.down('xs')]: {
      height: 470
    }
  },
  createHeader: customTheme.card.header,
  caption: {
    color: '#424850',
    marginBottom: theme.spacing(2),
    marginTop: theme.spacing(2),
  },
  labelCheckbox: {
    marginTop: theme.spacing(3),
  }
});

class CreateVolume extends Component {
  constructor(props) {
    super(props);
    this.state = {
      volume_count: 1,
      volume_name: 'vol',
      volume_suffix: 0,
      volume_size: 100,
      volume_description: '',
      volume_units: 'GB',
      form_valid: true,
      open: false,
      alert_description: '',
      maxbw: 0,
      maxiops: 0,
      stop_on_error_checkbox: false,
      mount_vol: true,
    };

    this.setUnit = this.setUnit.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.createVolumeInParent = this.createVolumeInParent.bind(this);
    this.handleClose = this.handleClose.bind(this);
  }


  componentDidMount() {

    // namespace to connect to the websocket for multi-volume creation
    const createVolSocketEndPoint = ":5000/create_vol";

    const createVolSocket = io(createVolSocketEndPoint, {
      transports: ['websocket'],
      query: {
        'x-access-token': localStorage.getItem('token')
      }
    });

    createVolSocket.on('connect', () => {
      console.log("connected to create volume socket")
    })
    // callback function for create multi-volume response
    createVolSocket.on("create_multi_volume", msg => {
      if (this.props.createVolumeButton === true) {
        let total_count = msg['total_count']
        let pass = msg['pass']
        let description = msg['description']
        let error_msg = 'Status: ' + pass + "/" + total_count + ' volume(s) created successfully'
        this.props.toggleCreateVolumeButton(false);
        this.props.showStorageAlert({
          alertType: 'info',
          alertTitle: 'Create Volume',
          errorMsg: error_msg,
          errorCode: description,
        });
        this.props.fetchVolumes()
      }
    });

  }

  setUnit(event) {
    this.setState({ volume_units: event.target.value });
  }

  handleChange(event) {

    const { name, value } = event.target;
    if (name === 'stop_on_error_checkbox') {
      this.setState({ stop_on_error_checkbox: !this.state.stop_on_error_checkbox });
    }
    else if (name === 'mount_vol_checkbox') {
      this.setState({ mount_vol: !this.state.mount_vol });
    }
    else
      this.setState({ [name]: value });
  }

  handleClose() {
    this.setState({ open: false });
  }

  createVolumeInParent() {
    let isError = true
    let errorDesc = ''
    if (this.state.volume_size.length === 0)
      errorDesc = "Please Enter Volume Size"
    else if (this.state.volume_size < 1)
      errorDesc = "Volume Size Should be greater than 0"
    else if (this.state.volume_name.length < 1)
      errorDesc = "Please Enter Volume Name"
    else if (this.state.volume_count.length === 0)
      errorDesc = "Please Enter Volume Count"
    else if (this.state.volume_count < 1)
      errorDesc = "Volume Count should be greater than 0"
    else if (this.state.volume_count > parseInt(this.props.maxVolumeCount,10))
      errorDesc = `Volume Count should not exceed ${ this.props.maxVolumeCount}`
    else if (this.state.volume_count > 1 && this.state.volume_suffix < 0)
      errorDesc = "Suffix Value cannot be negative"
    else if (this.state.volume_count > 1 && this.state.volume_suffix === null)
      errorDesc = "Please Enter Suffix Start Value"
    else if (this.state.maxbw.length === 0)
      errorDesc = "Please Enter Maximum Bandwidth (MB/s) "
    else if (this.state.maxiops.length === 0)
      errorDesc = "Please Enter Maximum IOPS (KIOPS)"
    else if (this.state.maxbw < 0)
      errorDesc = "Max Bandwidth cannot be negative"
    else if (this.state.maxiops < 0)
      errorDesc = "Maximum IOPS be negative"
    else if (this.state.maxiops > 0 && this.state.maxiops < 10)
      errorDesc = "Invalid value for Maximum IOPS"
    else
      isError = false

    if (isError === true) {
      this.setState({ open: true, alert_description: errorDesc });
      return;
    }

    this.props.createVolume({ ...this.state });

    this.setState({
      ...this.state,
      volume_count: 1,
      volume_name: 'vol',
      volume_suffix: 0,
      volume_size: 100,
      volume_description: '',
      volume_units: 'GB',
      maxbw: 0,
      maxiops: 0,
      stop_on_error_checkbox: false,
      mount_vol: true,
    });
  }

  render() {
    const { classes } = this.props;
    let volumeCountTitle
    if (this.props.volCount > 1)
      volumeCountTitle = `Specify the number of volumes to create. ${ this.props.volCount } volumes already exist. POS supports max ${ this.props.maxVolumeCount } volumes`
    else if (this.props.volCount === 1)
      volumeCountTitle = `Specify the number of volumes to create. ${ this.props.volCount } volume already exists. POS supports max ${ this.props.maxVolumeCount } volumes`
    else
      volumeCountTitle = `Specify the number of volumes to create. POS supports max ${ this.props.maxVolumeCount } volumes`

    return (
      <ThemeProvider theme={PageTheme}>
        <Paper className={classes.volumeCreatePaper}>
          <Grid item xs={12}>
            <Typography className={classes.createHeader}>
              Create Volume
            </Typography>
          </Grid>
          <form className={classes.formContainer}>
            <Grid item container xs={12} sm={6} justify="flex-start" className={classes.formControl}>
              <Tooltip
                title={volumeCountTitle}
                placement="bottom-start"
              >
                <FormControl className={classes.volumeName}>
                  <TextField
                    id="create-vol-count"
                    name="volume_count"
                    label="Volume Count"
                    type="number"
                    inputProps={{
                        min: 1,
                        max: this.props.maxVolumeCount,
                        'data-testid': "create-vol-count"
                    }}
                    value={this.state.volume_count}
                    onChange={this.handleChange}
                    required="true"
                  />
                </FormControl>
              </Tooltip>
            </Grid>
            <Grid item container xs={12} sm={6} justify="flex-end" className={classes.formControl}>
              <FormControl className={classes.volumeName}>
                <FormControlLabel
                  control={(
<Checkbox name="mount_vol_checkbox" color="primary" id="mount-vol-checkbox"
                      checked={this.state.mount_vol}
                      value="Mount Volume"
                      onChange={this.handleChange}
/>
)}

                  label="Mount Volume"
                  className={classes.labelCheckbox}
                />
              </FormControl>
            </Grid>
            <Grid item container xs={12}>
              <Typography variant="h4" component="h4" className={classes.caption} display="block">
                For Volume Count &gt; 1, please provide a seed in the Suffix Start Value field (e.g. 0,1)
              </Typography>
            </Grid>

            <Grid item container xs={12} sm={6} justify="flex-start" className={classes.formControl}>
              <FormControl className={classes.volumeName}>
                <TextField
                  id="create-vol-name"
                  label="Volume Name"
                  name="volume_name"
                  value={this.state.volume_name}
                  onChange={this.handleChange}
                  inputProps={{
                            'data-testid': "create-vol-name"
                  }}
                  required="true"
                />
              </FormControl>
            </Grid>
            <Grid item container xs={12} sm={6} justify="flex-end" className={classes.formControl}>
              <Tooltip
                title=" Min suffix value allowed is 0.
                        The suffix will be appended to the volume name to form the final volume name (e.g. vol_0, vol_1)"
                placement="bottom-start"
                disableFocusListener={this.state.volume_count < 2}
                disableHoverListener={this.state.volume_count < 2}
                disableTouchListener={this.state.volume_count < 2}
              >
                <FormControl className={classes.volumeName}>
                  <TextField
                    id="create-vol-suffix"
                    label="Suffix Start Value"
                    name="volume_suffix"
                    type="number"

                    InputProps={{ inputProps: { min: 0 } }}
                    value={this.state.volume_suffix}
                    onChange={this.handleChange}
                    disabled={this.state.volume_count < 2}
                  />
                </FormControl>
              </Tooltip>
            </Grid>
            <Grid item container xs={12} sm={6} justify="flex-start" className={classes.formControl}>
              <FormControl className={classes.unitText}>
                <TextField
                  id="create-vol-size"
                  label="Volume Size"
                  name="volume_size"
                  value={this.state.volume_size}
                  onChange={this.handleChange}
                  type="number"
                  inputProps={{
                            'data-testid': "create-vol-size"
                  }}
                  required="true"
                />
              </FormControl>
              <FormControl className={classes.volumeUnit}>
                <Select
                  value={this.state.volume_units}
                  onChange={this.setUnit}
                  inputProps={{
                    name: 'Volume Unit',
                    id: 'vol_unit'
                  }}
                  SelectDisplayProps={{
                    'data-testid': 'volume-unit'
                  }}
                  className={classes.unitSelect}
                >
                  <MenuItem value="GB">GB</MenuItem>
                  <MenuItem value="TB">TB</MenuItem>
                </Select>
              </FormControl>
            </Grid>


            <Grid item container xs={12} sm={6} justify="flex-end" className={classes.formControl}>
              <Tooltip title="Min Value 10. 0 means max" placement="bottom-start">
                <FormControl className={classes.volumeName}>
                  <TextField
                    id="create-vol-maxiops"
                    label="Maximum IOPS (KIOPS)"
                    name="maxiops"
                    value={this.state.maxiops}
                    onChange={this.handleChange}
                    type="number"
                    // placeholder="Min Value 10. 0 means max"
                    inputProps={{ min: 0, 'data-testid': "create-vol-max-iops" }}
                    required="true"
                  />
                </FormControl>
              </Tooltip>
            </Grid>
            <Grid item container xs={12} sm={6} justify="flex-start" className={classes.formControl}>
              <Tooltip title="0 means max" placement="bottom-start">
                <FormControl className={classes.volumeName}>
                  <TextField
                    id="create-vol-maxbw"
                    label="Maximum Bandwidth (MB/s)"
                    name="maxbw"
                    value={this.state.maxbw}
                    onChange={this.handleChange}
                    type="number"
                    // placeholder="0 means max"
                    inputProps={{ min: 0, 'data-testid': "create-vol-max-bw" }}
                    required="true"
                  />
                </FormControl>
              </Tooltip>
            </Grid>


            <Grid item container xs={12} sm={6} justify="flex-end" className={classes.formControl}>
              <FormControl className={classes.volumeName}>
                <Tooltip
                  title="Do you want to proceed with subsequent volume creation in case an error occurs or abort the remaining process?"
                  placement="top-start"
                  disableFocusListener={this.state.volume_count < 2}
                  disableHoverListener={this.state.volume_count < 2}
                  disableTouchListener={this.state.volume_count < 2}
                >
                  <FormControlLabel
                    disabled={this.state.volume_count < 2}
                    control={(
<Checkbox name="stop_on_error_checkbox" color="primary" id="create-vol-stop-on-error-checkbox"
                        checked={this.state.stop_on_error_checkbox}
                        value="Stop on error"
                        onChange={this.handleChange}
/>
)}
                    label="Stop Multi-Volume Creation on Error"
                    className={classes.labelCheckbox}
                  />
                </Tooltip>
              </FormControl>
            </Grid>
            {/* <Grid item container xs={12} sm={6} justify="flex-end" className={classes.formControl}>
              <FormControl className={classes.volumeName}>
                  <FormControlLabel
                    control={
                      <Checkbox name="mount_vol_checkbox" color="primary" id="mount-vol-checkbox"
                        checked={this.state.mount_vol}
                        value="Mount Volume"
                        onChange={this.handleChange}
                      />
                    }

                    label="Mount Volume"
                    className={classes.labelCheckbox}
                  />
              </FormControl>
            </Grid> */}
            <Grid item container xs={12} display="flex" justify="flex-start" className={`${classes.volBtnContainer} ${classes.formControl}`}>
              <Tooltip
                title="Volume Creation is in progress. Please wait for sometime."
                placement="right-start"
                open={this.props.createVolumeButton}
              >
                <Button
                  onClick={this.createVolumeInParent}
                  variant="contained"
                  color="primary"
                  data-testid="createvolume-btn"
                  className={classes.button}
                  disabled={this.props.createVolumeButton}
                >
                  Create Volume
                </Button>
              </Tooltip>
            </Grid>
          </form>
          <AlertDialog
            title="Create Volume"
            description={this.state.alert_description}
            type="alert"
            open={this.state.open}
            handleClose={this.handleClose}
          />
        </Paper>
      </ThemeProvider>
    );
  }
}

CreateVolume.propTypes = {
  createVolume: PropTypes.func.isRequired
}

const mapStateToProps = state => {
  return {
    createVolumeButton: state.storageReducer.createVolumeButton,
  };
}

const mapDispatchToProps = dispatch => {
  return {
    toggleCreateVolumeButton: (flag) => dispatch(actionCreators.toggleCreateVolumeButton(flag)),
    showStorageAlert: (alertParams) => dispatch(actionCreators.showStorageAlert(alertParams)),
    fetchVolumes: () => dispatch({ type: actionTypes.SAGA_FETCH_VOLUMES }),
  };
}


export default withStyles(styles)((connect(mapStateToProps, mapDispatchToProps))(((CreateVolume))));
