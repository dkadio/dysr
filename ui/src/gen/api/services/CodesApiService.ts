/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CreateCode_FmInput } from '../models/CreateCode_FmInput';
import type { ModelsHealthcheck } from '../models/ModelsHealthcheck';
import type { ModelsUserCode } from '../models/ModelsUserCode';
import type { UpdateCode_FmInput } from '../models/UpdateCode_FmInput';
import type { CancelablePromise } from '../core/CancelablePromise';
import { request as __request } from '../core/request';

export class CodesApiService {
  /**
   * Get all User Codes.
   * @returns ModelsUserCode OK
   * @throws ApiError
   */
  public static getCodesFm(): CancelablePromise<Array<ModelsUserCode>> {
    return __request({
      method: 'GET',
      path: `/api/v1/codes`,
      errors: {
        404: `Not Found`,
        500: `Internal Server Error`
      }
    });
  }

  /**
   * Update a pet.
   * @param id
   * @param requestBody
   * @returns ModelsUserCode OK
   * @throws ApiError
   */
  public static createCodeFm(
    id: string,
    requestBody?: CreateCode_FmInput
  ): CancelablePromise<ModelsUserCode> {
    return __request({
      method: 'POST',
      path: `/api/v1/codes`,
      body: requestBody,
      mediaType: 'application/json',
      errors: {
        500: `Internal Server Error`
      }
    });
  }

  /**
   * Get all User Codes.
   * @param id
   * @returns ModelsUserCode OK
   * @throws ApiError
   */
  public static getCodeFm(id: string): CancelablePromise<ModelsUserCode> {
    return __request({
      method: 'GET',
      path: `/api/v1/codes/${id}`,
      errors: {
        404: `Not Found`,
        500: `Internal Server Error`
      }
    });
  }

  /**
   * Update a pet.
   * @param id
   * @param requestBody
   * @returns ModelsUserCode OK
   * @throws ApiError
   */
  public static updateCodeFm(
    id: string,
    requestBody?: UpdateCode_FmInput
  ): CancelablePromise<ModelsUserCode> {
    return __request({
      method: 'PUT',
      path: `/api/v1/codes/${id}`,
      body: requestBody,
      mediaType: 'application/json',
      errors: {
        404: `Not Found`,
        500: `Internal Server Error`
      }
    });
  }

  /**
   * Deletes a UserCode.
   * @param id
   * @returns any OK
   * @throws ApiError
   */
  public static deleteCodeFm(id: string): CancelablePromise<any> {
    return __request({
      method: 'DELETE',
      path: `/api/v1/codes/${id}`,
      errors: {
        404: `Not Found`,
        500: `Internal Server Error`
      }
    });
  }

  /**
   * Checks API is healthy.
   * @returns ModelsHealthcheck OK
   * @throws ApiError
   */
  public static healthcheck(): CancelablePromise<ModelsHealthcheck> {
    return __request({
      method: 'GET',
      path: `/api/v1/healthcheck`,
      errors: {
        500: `Internal Server Error`
      }
    });
  }
}
