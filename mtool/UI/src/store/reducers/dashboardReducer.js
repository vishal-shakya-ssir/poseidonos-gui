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


DESCRIPTION: <Contains reducer function for dashboard page> *
@NAME : dashboardReducer.js
@AUTHORS: Jay Hitesh Sanghavi 
@Version : 1.0 *
@REVISION HISTORY
[03/06/2019] [Jay] : Prototyping..........////////////////////
*/

import * as actionTypes from "../actions/actionTypes"

const initialState = {
    volumes: [],
    arrayVolumes: [],
    alerts: [],
    selectedArray: 'all',
    ibofs: ['None'],
    readIOPS: 0,
    writeIOPS: 0,
    readBW: 0,
    writeBW: 0,
    latency: 0,
    fetchingAlerts: false,
    ip: '0.0.0.0',
    mac: 'NA',
    host: '',
    arraySize: 0,
    arrayVols: {}
}


const dashboardReducer = (state = initialState, action) => {
    switch (action.type) {
        case actionTypes.ENABLE_FETCHING_ALERTS:
            return {
                ...state,
                fetchingAlerts: action.fetchingAlerts,
            };
        case actionTypes.SELECT_ARRAY: {
            let arrayVolumes = [];
            if(action.array === "all") {
                arrayVolumes = [...state.volumes]
            } else {
                state.volumes.forEach((volume) => {
                    if(volume.array === action.array) {
                        arrayVolumes.push({
                            ...volume
                        });
                    }
                });
            }
            return {
                ...state,
                selectedArray: action.array,
                arrayVolumes
            }
        }
        case actionTypes.FETCH_VOLUME_INFO:{
            const volumes = [];
            const arrayVols = {};
            let arrayVolumes = [];
            Object.keys(action.volumes).forEach((array) => {
                arrayVols[array] = action.volumes[array].length;
                action.volumes[array].forEach((vol) => {
                    volumes.push({
                        ...vol,
                        array
                    });
                    if(state.selectedArray !== "all" && array === state.selectedArray) {
                        arrayVolumes.push({
                            ...volumes[volumes.length - 1]
                        })
                    }
                });
            });
            if(state.selectedArray === "all") {
                arrayVolumes = [...volumes];
            }
            return {
                ...state,
                volumes,
                arrayVols,
                arrayVolumes
            };
        }
        case actionTypes.FETCH_ALERTS_INFO:
            return {
                ...state,
                alerts: action.alerts,
            };
        case actionTypes.FETCH_STORAGE_INFO:
            return {
                ...state,
                arraySize: action.arraySize
            };
        case actionTypes.FETCH_PERFORMANCE_INFO:
            return {
                ...state,
                readIOPS: action.readIOPS,
                writeIOPS: action.writeIOPS,
                readBW: action.readBW,
                writeBW: action.writeBW,
                latency: action.latency
            };
        case actionTypes.FETCH_IPANDMAC_INFO:
            return {
                ...state,
                ip: action.ip,
                mac: action.mac,
                host: action.host
            };
        default:
            return state;
    }
};

export default dashboardReducer;