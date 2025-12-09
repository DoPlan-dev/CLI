
import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DataService } from '../services/data.service';

@Component({
  selector: 'app-achievements',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="space-y-8">
      
      <!-- Hero Section -->
      <div class="bg-white rounded-2xl border border-gray-200 p-8 relative overflow-hidden flex flex-col md:flex-row items-center shadow-sm">
        <!-- Background Decor -->
        <div class="absolute right-0 top-0 h-full w-1/3 bg-gradient-to-l from-blue-50 to-transparent pointer-events-none"></div>
        
        <div class="relative z-10 flex flex-col md:flex-row items-center w-full">
          <!-- Level Circle -->
          <div class="relative w-32 h-32 flex-shrink-0 mb-6 md:mb-0 md:mr-8 group cursor-pointer">
            <svg class="w-full h-full transform -rotate-90">
              <circle cx="64" cy="64" r="58" stroke="#f1f5f9" stroke-width="8" fill="none" />
              <circle cx="64" cy="64" r="58" stroke="#206bc4" stroke-width="8" fill="none" stroke-dasharray="364" stroke-dashoffset="90" stroke-linecap="round" class="transition-all group-hover:stroke-[#1a5c98]" />
            </svg>
            <div class="absolute inset-0 flex flex-col items-center justify-center">
              <span class="text-3xl font-bold text-gray-800">Lvl 5</span>
              <span class="text-xs text-gray-500">XP 750/1000</span>
            </div>
          </div>

          <div class="text-center md:text-left flex-1">
            <h2 class="text-3xl font-bold text-gray-800 mb-2">Achievement Master</h2>
            <p class="text-gray-500 max-w-lg mb-4">You're on a roll! Complete 3 more tasks to unlock the "Productivity Beast" badge and reach Level 6.</p>
            <div class="flex flex-wrap gap-2 justify-center md:justify-start">
               <span class="px-3 py-1 bg-yellow-100 text-yellow-700 border border-yellow-200 rounded-full text-xs font-semibold flex items-center gap-1">
                  🔥 5 Day Streak
               </span>
               <span class="px-3 py-1 bg-purple-100 text-purple-700 border border-purple-200 rounded-full text-xs font-semibold flex items-center gap-1">
                  🏆 Top 5% Contributor
               </span>
            </div>
          </div>
          
          <!-- Illustration -->
          <div class="hidden lg:block w-48 h-48 ml-8">
              <img src="images/loading.avif" class="w-full h-full object-contain" alt="Achievement Unlocked">
          </div>
        </div>
      </div>

      <!-- Achievements Grid -->
      <div>
        <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-bold text-gray-800">Badges & Milestones</h3>
            <div class="flex gap-2">
                <button class="text-sm px-3 py-1 rounded-full bg-blue-100 text-[#206bc4] font-medium">All</button>
                <button class="text-sm px-3 py-1 rounded-full text-gray-500 hover:bg-gray-100">Locked</button>
                <button class="text-sm px-3 py-1 rounded-full text-gray-500 hover:bg-gray-100">Unlocked</button>
            </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          @for (ach of dataService.achievements(); track ach.id) {
            <div class="bg-white rounded-xl border p-6 flex flex-col items-center text-center transition-all duration-300 group hover:-translate-y-1 hover:shadow-lg"
                 [class.border-gray-200]="!ach.unlocked"
                 [class.border-blue-200]="ach.unlocked"
                 [class.opacity-60]="!ach.unlocked"
                 [class.grayscale]="!ach.unlocked"
                 (click)="dataService.toggleAchievement(ach.id)">
              
              <div class="w-16 h-16 rounded-full mb-4 flex items-center justify-center text-2xl relative transition-all duration-500"
                   [class.bg-gray-100]="!ach.unlocked"
                   [class.bg-blue-50]="ach.unlocked"
                   [class.text-[#206bc4]]="ach.unlocked"
                   [class.scale-110]="ach.unlocked">
                 
                 <!-- Icons -->
                 @if (ach.icon === 'sun') { ☀️ }
                 @else if (ach.icon === 'bug') { 🐛 }
                 @else if (ach.icon === 'flame') { 🔥 }
                 @else if (ach.icon === 'users') { 👥 }
              </div>

              <h4 class="text-gray-800 font-bold mb-1">{{ ach.title }}</h4>
              <p class="text-sm text-gray-500 mb-4 h-10">{{ ach.description }}</p>

              <div class="w-full bg-gray-100 h-2 rounded-full overflow-hidden mt-auto">
                 <div class="bg-[#206bc4] h-full transition-all duration-1000" [style.width.%]="ach.progress"></div>
              </div>
              <span class="text-xs text-gray-400 mt-2 block">{{ ach.progress }}% Complete</span>
            </div>
          }
        </div>
      </div>
    </div>
  `
})
export class AchievementsComponent {
  dataService = inject(DataService);
}
