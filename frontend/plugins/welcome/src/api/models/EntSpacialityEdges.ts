/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntOrgan,
    EntOrganFromJSON,
    EntOrganFromJSONTyped,
    EntOrganToJSON,
    EntPhysician,
    EntPhysicianFromJSON,
    EntPhysicianFromJSONTyped,
    EntPhysicianToJSON,
    EntTypeDisease,
    EntTypeDiseaseFromJSON,
    EntTypeDiseaseFromJSONTyped,
    EntTypeDiseaseToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSpacialityEdges
 */
export interface EntSpacialityEdges {
    /**
     * 
     * @type {EntOrgan}
     * @memberof EntSpacialityEdges
     */
    ogan?: EntOrgan;
    /**
     * 
     * @type {EntTypeDisease}
     * @memberof EntSpacialityEdges
     */
    type?: EntTypeDisease;
    /**
     * 
     * @type {EntPhysician}
     * @memberof EntSpacialityEdges
     */
    user?: EntPhysician;
}

export function EntSpacialityEdgesFromJSON(json: any): EntSpacialityEdges {
    return EntSpacialityEdgesFromJSONTyped(json, false);
}

export function EntSpacialityEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSpacialityEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ogan': !exists(json, 'Ogan') ? undefined : EntOrganFromJSON(json['Ogan']),
        'type': !exists(json, 'Type') ? undefined : EntTypeDiseaseFromJSON(json['Type']),
        'user': !exists(json, 'User') ? undefined : EntPhysicianFromJSON(json['User']),
    };
}

export function EntSpacialityEdgesToJSON(value?: EntSpacialityEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ogan': EntOrganToJSON(value.ogan),
        'type': EntTypeDiseaseToJSON(value.type),
        'user': EntPhysicianToJSON(value.user),
    };
}

