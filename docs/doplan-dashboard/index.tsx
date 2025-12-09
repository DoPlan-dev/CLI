
import { bootstrapApplication } from '@angular/platform-browser';
import { provideZonelessChangeDetection } from '@angular/core';
import { provideRouter, withHashLocation, Routes } from '@angular/router';
import { AppComponent } from './src/app.component';
import { DashboardComponent } from './src/components/dashboard.component';
import { PlanComponent } from './src/components/plan.component';
import { MeetingsComponent } from './src/components/meetings.component';
import { AchievementsComponent } from './src/components/achievements.component';
import { SettingsComponent } from './src/components/settings.component';

const routes: Routes = [
  { path: '', component: DashboardComponent },
  { path: 'plan', component: PlanComponent },
  { path: 'meetings', component: MeetingsComponent },
  { path: 'achievements', component: AchievementsComponent },
  { path: 'settings', component: SettingsComponent },
  { path: '**', redirectTo: '' }
];

bootstrapApplication(AppComponent, {
  providers: [
    provideZonelessChangeDetection(),
    provideRouter(routes, withHashLocation())
  ]
}).catch(err => console.error(err));

// AI Studio always uses an `index.tsx` file for all project types.
