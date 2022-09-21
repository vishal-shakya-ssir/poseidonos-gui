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

import { all } from 'redux-saga/effects';
import headerWatcher from "./headerSaga"
import {dashboardWatcher} from "./dashboardSaga"
import {configurationsettingWatcher} from "./configurationsettingSaga"
import {logManagementWatcher} from "./logManagementSaga"
import {alertManagementContainerWatcher} from "./alertManagementSaga"
import {alertManagementAlertTableWatcher} from "./alertManagementAlertTableSaga"
import {alertManagementAddNewAlertsWatcher} from "./alertManagementAddNewAlertsSaga"
import {authenticationWatcher} from "./authenticationSaga"
import {userManagementContainerWatcher} from "./userManagementSaga"
import {userManagementUserTableWatcher} from "./userManagementUserTableSaga"
import {userManagementAddNewUsersWatcher} from "./userManagementAddNewUsersSaga"
import storageWatcher from "./storageSaga";
import performanceWatcher from './performanceSaga';
import {hardwareOverviewWatcher} from './hardwareOverviewSaga'
import {hardwareSensorWatcher} from './hardwareSensorSaga'
import {hardwareHealthWatcher} from './hardwareHealthSaga'
import {hardwarePowerManagementWatcher} from './hardwarePowerManagementSaga'
import {BMCAuthenticationWatcher} from './BMCAuthenticationSaga'
import subsystemWatcher from './subsystemSaga';
import { telemetryWatcher } from './telemetrySaga';

export default function* rootSaga() {
    yield all([
        headerWatcher(),
        dashboardWatcher(),
        configurationsettingWatcher(),
        logManagementWatcher(),
        alertManagementContainerWatcher(),
        alertManagementAlertTableWatcher(),
        alertManagementAddNewAlertsWatcher(),
        storageWatcher(),
        subsystemWatcher(),
        performanceWatcher(),
        authenticationWatcher(),
        userManagementContainerWatcher(),
        userManagementUserTableWatcher(),
        userManagementAddNewUsersWatcher(),
        hardwareOverviewWatcher(),
        hardwareSensorWatcher(),
        hardwareHealthWatcher(),
        hardwarePowerManagementWatcher(),
        BMCAuthenticationWatcher(),
        telemetryWatcher(),
    ]);
}
