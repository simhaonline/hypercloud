import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ClarityModule } from '@clr/angular';
import { HttpModule } from '@angular/http';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './components/app/app.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ClusterDetailComponent } from './components/cluster/cluster-detail.component';
import { ServiceListComponent } from './components/service/service-list.component';
import { ServiceCreateComponent } from './components/service/service-create.component';
import { ServiceDetailComponent } from './components/service/service-detail.component';
import { ServiceDetailOverviewComponent } from './components/service/service-detail-overview.component';
import { ServiceDetailMetricsComponent } from './components/service/service-detail-metrics.component';
import { ServiceDetailSettingsComponent } from './components/service/service-detail-settings.component';
import { ApplicationNameValidatorDirective } from './directives/application-name-validator.directive';
import { StackListComponent } from './components/stack/stack-list.component';
import { StackDetailComponent } from './components/stack/stack-detail.component';

import {
    ApiService,
    StackService,
    ServiceService,
    DockerService,
    EventService
} from './services';

import {
    TruncatePipe,
    SizePipe,
    ImagePipe,
    ContainerPortPipe,
    StackNamePipe,
    ServiceNamePipe
} from './pipes';

import { ServiceResolver } from './resolvers/service.resolver';

@NgModule({
    declarations: [
        AppComponent,
        DashboardComponent,
        ClusterDetailComponent,
        ServiceListComponent,
        ServiceCreateComponent,
        ServiceDetailComponent,
        ServiceDetailOverviewComponent,
        ServiceDetailMetricsComponent,
        ServiceDetailSettingsComponent,
        ApplicationNameValidatorDirective,
        StackListComponent,
        StackDetailComponent,
        TruncatePipe,
        SizePipe,
        ImagePipe,
        ContainerPortPipe,
        StackNamePipe,
        ServiceNamePipe
    ],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        FormsModule,
        AppRoutingModule,
        ClarityModule,
        HttpModule
    ],
    providers: [
        StackNamePipe,
        ApiService,
        StackService,
        ServiceService,
        DockerService,
        EventService,
        ServiceResolver
    ],
    bootstrap: [AppComponent]
})
export class AppModule { }
