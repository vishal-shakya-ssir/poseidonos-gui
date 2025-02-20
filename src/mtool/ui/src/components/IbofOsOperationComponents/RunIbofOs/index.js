/*
 *   BSD LICENSE
 *   Copyright (c) 2021 Samsung Electronics Corporation
 *   All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in
 *       the documentation and/or other materials provided with the
 *       distribution.
 *     * Neither the name of Samsung Electronics Corporation nor the names of its
 *       contributors may be used to endorse or promote products derived
 *       from this software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

import React, { useState } from 'react';
import {
  Button,
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Select,
  Typography
} from '@material-ui/core';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import ErrorIcon from '@material-ui/icons/Error';
import './RunIbofOs.css';
import LinearProgressBarComponent from '../LinearProgressBarComponent';
import { customTheme } from '../../../theme';

const RunIbofOs = props => {
  
  const [propertySelect, setPropertySelect] = useState("");
  const selectProperty = (event) => {
    setPropertySelect(event.target.value);
  }
  const setProperty = () => {
    props.setProperty(propertySelect);
  }

  return (
    <div className="RunIbofOs-Outer-Box">
      <Typography className="IBOFOSMainHeader" variant="h1">Run/Shutdown Poseidon OS</Typography>
      <div>
        <span className="IBOFOSNoteMessage">
          <span style={{ float: 'left', marginRight: '2px' }}>
            Status -
          </span>
          <span
            style={{ float: 'left', color: (props.status ? customTheme.palette.success.dark : customTheme.palette.error.dark), fontWeight: '600' }}
          >
            {props.status ? "Running" : props.OS_Running_Status}
          </span>
          {props.OS_Running_Status !== 'Not Running' && localStorage.getItem('Rebuilding_Value') ? (
            <LinearProgressBarComponent
              percent={localStorage.getItem('Rebuilding_Value')}
            />
          ) : null}
        </span>
        <div className="IBOFOSCommandWrapper">
          <span className="IBOFOSStartLabel">
            <span style={{ float: 'left' }}>
              Start/Stop Poseidon OS
            </span>
          </span>
          <div>
            <div className="IBOFOSButtonClass">
              <Button
                variant="contained"
                color="secondary"
                title="Force Start Poseidon OS"
                data-testid="startButton"
                disabled={props.status}
                onClick={() => {
                  props.openAlert('Start');
                }}
              >
                Start
              </Button>
            </div>
            <div className="IBOFOSButtonClass">
              <Button
                variant="outlined"
                color="secondary"
                title="Force Stop Poseidon OS"
                data-testid="stopButton"
                disabled={!props.status}
                onClick={() => {
                  props.openAlert('Stop');
                }}
              >
                Stop
              </Button>
            </div>
          </div>
        </div>
      </div>
      <Grid container xs={12} className="IBOFOSSetPropertyContainer" alignItems="baseline">
        <FormControl>
          <InputLabel htmlFor="vol_unit">Rebuild Perf Impact</InputLabel>
          <Select
            value={propertySelect}
            onChange={selectProperty}
            disabled={props.OS_Running_Status === 'Not Running'}
            inputProps={{
              name: "Volume Unit",
              id: "vol_unit",
              "data-testid": "set-property-input",
            }}
            SelectDisplayProps={{
              "data-testid": "set-property-select",
            }}
            className="IBOFOSPropertySelect"
          >
            <MenuItem value="high" data-testid="high">
              High
            </MenuItem>
            <MenuItem value="medium" data-testid="medium">
              Medium
            </MenuItem>
            <MenuItem value="low" data-testid="lowest">
              Low
            </MenuItem>
          </Select>
        </FormControl>
        <Button
          variant="contained"
          color="secondary"
          size="small"
          title="Set Rebuild Performance Impact"
          data-testid="setPropertyButton"
          className="IBOFOSPropertyButton"
          onClick={setProperty}
          disabled={props.OS_Running_Status === 'Not Running' || propertySelect === ""}
        >
          Set Property
        </Button>
        {props.OS_Running_Status !== 'Not Running' ? (
          <Typography className="IBOFOSImpactText"> Current Impact: <span className="IBOFOSUppercase">{props.property}</span></Typography>
        ) : null}
      </Grid>
      <Grid className="IBOFOSStatusLabel">
        <span>Response</span>
        <Grid className="IBOFOSResponse">
          {props.responsefromos && props.responsefromos.map((response) => (
            <Typography component="div" className="IBOFOSStatusText">{response.code === 0 ? (
              <CheckCircleIcon className="IBOFOSStatusIcon" style={{ color: customTheme.palette.success.dark }} />
            ) : null}{response.code &&
              response.code !== 200 ? (
              <ErrorIcon className="IBOFOSStatusIcon" style={{ color: customTheme.palette.error.dark }} />
            ) : null} {response.description}
            </Typography>
          ))}
        </Grid>
      </Grid>
    </div>
  );
};
export default RunIbofOs;
