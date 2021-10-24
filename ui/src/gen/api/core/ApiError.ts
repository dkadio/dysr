/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ApiResult } from './ApiResult';

export class ApiError extends Error {
  url: string;
  status: number;
  statusText: string;
  body: any;

  constructor(response: ApiResult, message: string) {
    super(message);

    this.name = 'ApiError';
    this.url = response.url;
    this.status = response.status;
    this.statusText = response.statusText;
    this.body = response.body;
  }
}
