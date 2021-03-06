import { Component, OnInit } from '@angular/core';
import { Service } from '../../entities';
import { ServiceService } from '../../services';

@Component({
    selector: 'app-service-list',
    templateUrl: './service-list.component.html',
    styleUrls: ['./service-list.component.less']
})
export class ServiceListComponent implements OnInit {
    public services: Service[] = [];

    constructor(private serviceService: ServiceService) {
    }

    ngOnInit() {
        this.serviceService.getServices().subscribe(services => this.services = services);
    }
}
