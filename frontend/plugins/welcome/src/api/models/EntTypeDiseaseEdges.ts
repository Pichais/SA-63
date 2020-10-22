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
    EntSpaciality,
    EntSpacialityFromJSON,
    EntSpacialityFromJSONTyped,
    EntSpacialityToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTypeDiseaseEdges
 */
export interface EntTypeDiseaseEdges {
    /**
     * Fortype holds the value of the fortype edge.
     * @type {Array<EntSpaciality>}
     * @memberof EntTypeDiseaseEdges
     */
    fortype?: Array<EntSpaciality>;
}

export function EntTypeDiseaseEdgesFromJSON(json: any): EntTypeDiseaseEdges {
    return EntTypeDiseaseEdgesFromJSONTyped(json, false);
}

export function EntTypeDiseaseEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTypeDiseaseEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fortype': !exists(json, 'fortype') ? undefined : ((json['fortype'] as Array<any>).map(EntSpacialityFromJSON)),
    };
}

export function EntTypeDiseaseEdgesToJSON(value?: EntTypeDiseaseEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fortype': value.fortype === undefined ? undefined : ((value.fortype as Array<any>).map(EntSpacialityToJSON)),
    };
}

