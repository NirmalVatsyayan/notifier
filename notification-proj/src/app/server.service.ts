/**
 * Created by Nirmal.Vatsyayan on 6/26/2017.
 */
import { Injectable } from '@angular/core';
import { Headers, Http, Response } from '@angular/http';
import 'rxjs/Rx';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class ServerService {
  constructor(private http: Http) {}
  storeNotification(notification_obj) {
    console.log(notification_obj)
    const headers = new Headers({'Content-Type': 'application/json'});
    return this.http.post('http://127.0.0.1:8080/notification',
      notification_obj,
      {headers: headers});
  }
}
