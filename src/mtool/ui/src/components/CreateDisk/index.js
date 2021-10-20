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
 *     * Neither the name of Intel Corporation nor the names of its
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

import React from 'react';
import { Button, FormControl, InputLabel, MenuItem, Select, TextField, withStyles } from '@material-ui/core';
import { connect } from 'react-redux';
import * as actionTypes from '../../store/actions/actionTypes';

const styles = (theme) => ({
    createDiskForm: {
        display: 'flex',
        flexDirection: 'column'
    },
    formItem: {
        margin: theme.spacing(1)
    }
});

const CreateDisk = (props) => {
    const { classes } = props;

    const createDiskSubmit = (event) => {
        event.preventDefault();
        props.Create_Disk({
            name: event.target.name.value,
            block_size: parseInt(event.target.block_size.value, 10),
            num_blocks: parseInt(event.target.num_blocks.value, 10),
            numa: parseInt(event.target.numa.value, 10),
            dev_type: event.target.dev_type.value
        });
        event.target.reset();
        props.cleanup();
    }

    return (
        <form className={classes.createDiskForm} onSubmit={createDiskSubmit}>
            <FormControl
                className={classes.formItem}
                size="small"
            >
                <TextField
                    label="Disk Name"
                    name="name"
                />
            </FormControl>
            <FormControl
              className={classes.formItem}
            >
                <TextField
                    label="Block Size"
                    name="block_size"
                    type="number"
                />
            </FormControl>
            <FormControl
                className={classes.formItem}
                size="small"
            >
                <TextField
                    label="Number of Blocks"
                    name="num_blocks"
                    type="number"
                />
            </FormControl>
            <FormControl
                className={classes.formItem}
                size="small"
            >
                <TextField
                    label="NUMA"
                    name="numa"
                    defaultValue={0}
                    type="number"
                />
            </FormControl>
            <FormControl
                size="small"
                className={classes.formItem}
            >
                <InputLabel htmlFor="disktype">Disk Type</InputLabel>
                <Select
                  inputProps={{
                      id: "disktype"
                  }}
                  name="dev_type"
                  defaultValue="uram"
                >
                  <MenuItem value="uram" data-testid="uram">
                    URAM
                  </MenuItem>
                  <MenuItem value="pram" data-testid="pram">
                    PRAM
                  </MenuItem>
                </Select>
            </FormControl>
            <Button
                className={classes.formItem}
                variant="contained"
                color="primary"
                type="submit"
            >
                Create Disk
            </Button>
        </form>
    );
}

const mapStateToProps = () => ({});

const mapDispatchToProps = (dispatch) => ({
    Create_Disk: (payload) => dispatch({type: actionTypes.SAGA_CREATE_DISK, payload})
});

export default connect(mapStateToProps, mapDispatchToProps)(withStyles(styles)(CreateDisk));