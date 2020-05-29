import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders, HttpResponse} from "@angular/common/http";
import {Observable, throwError} from "rxjs";
import {catchError} from "rxjs/operators";

import {Resource} from "../_models/resource";
import {Stage} from "../_models/stage";
import {ProjectResult} from "../_models/project-result";
import {ServiceResult} from "../_models/service-result";
import {EventResult} from "../_models/event-result";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private _baseUrl: string;
  private headers: HttpHeaders = new HttpHeaders({'Content-Type': 'application/json'});

  set baseUrl(value: string) {
    this._baseUrl = value;
  }

  constructor(private http: HttpClient) {
    this._baseUrl = `/api`;
  }

  public getVersion(): Observable<Object> {
    let url = `/api/`;
    return this.http
      .get<Object>(url, { headers: this.headers })
      .pipe(catchError(this.handleError<Object>('getVersion')));
  }

  public getProjects(): Observable<ProjectResult> {
    let url = `${this._baseUrl}/configuration-service/v1/project?disableUpstreamSync=true`;
    return this.http
      .get<ProjectResult>(url, { headers: this.headers })
      .pipe(catchError(this.handleError<ProjectResult>('getProjects')));
  }

  public getServices(projectName, stageName): Observable<ServiceResult> {
    let url = `${this._baseUrl}/configuration-service/v1/project/${projectName}/stage/${stageName}/service`;
    return this.http
      .get<ServiceResult>(url, { headers: this.headers })
      .pipe(catchError(this.handleError<ServiceResult>('getServices')));
  }

  public getRoots(projectName: string, serviceName: string, fromTime?: string): Observable<HttpResponse<EventResult>> {
    let url = `${this._baseUrl}/mongodb-datastore/event?root=true&pageSize=20&project=${projectName}&service=${serviceName}`;
    if(fromTime)
      url += `&fromTime=${fromTime}`;
    return this.http
      .get<EventResult>(url, { headers: this.headers, observe: 'response' })
      .pipe(catchError(this.handleError<HttpResponse<EventResult>>('getRoots')));
  }

  public getTraces(contextId: string, projectName?: string, fromTime?: string): Observable<HttpResponse<EventResult>> {
    let url = `${this._baseUrl}/mongodb-datastore/event?pageSize=100&keptnContext=${contextId}`;
    if(projectName)
      url += `&project=${projectName}`;
    if(fromTime)
      url += `&fromTime=${fromTime}`;
    return this.http
      .get<EventResult>(url, { headers: this.headers, observe: 'response' })
      .pipe(catchError(this.handleError<HttpResponse<EventResult>>('getTraces')));
  }

  public getEvaluationResults(projectName: string, serviceName: string, stageName: string, source: string, fromTime?: string) {
    let url = `${this._baseUrl}/mongodb-datastore/event?type=sh.keptn.events.evaluation-done&project=${projectName}&service=${serviceName}&stage=${stageName}&source=${source}&pageSize=50`;
    if(fromTime)
      url += `&fromTime=${fromTime}`;
    return this.http
      .get<EventResult>(url, { headers: this.headers })
      .pipe(catchError(this.handleError<EventResult>('getEvaluationResults')));
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      // TODO: handel error and show to the user?!
      this.log(`${operation} failed: ${error.message}`);
      return throwError(error);
    };
  }

  private log(message: string) {
    console.log(message);
  }

}
