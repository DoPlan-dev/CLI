import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DataService } from '../services/data.service';

@Component({
  selector: 'app-meetings',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="space-y-6">
      
      <div class="flex justify-between items-center">
        <h2 class="text-3xl font-bold text-gray-800">Meetings</h2>
        <button class="px-4 py-2 bg-[#206bc4] text-white font-bold rounded-lg text-sm shadow-sm hover:bg-[#1a5c98] transition-all">
            + Schedule Meeting
        </button>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Upcoming Schedule List -->
        <div class="lg:col-span-2 space-y-4">
          <h3 class="text-lg font-semibold text-gray-700">Upcoming</h3>
          
          @for (meeting of dataService.meetings(); track meeting.id) {
             <div class="bg-white p-5 rounded-xl border border-gray-200 flex items-center hover:border-[#206bc4] hover:shadow-md transition-all group shadow-sm">
                <!-- Date Badge -->
                <div class="flex-shrink-0 w-16 h-16 bg-gray-50 rounded-lg border border-gray-200 flex flex-col items-center justify-center text-center mr-5 group-hover:bg-[#206bc4]/5 group-hover:border-[#206bc4]/20 transition-all">
                  <span class="text-xs text-gray-500 uppercase font-bold">{{ formatDateMonth(meeting.date) }}</span>
                  <span class="text-xl font-bold text-gray-800">{{ formatDateDay(meeting.date) }}</span>
                </div>

                <!-- Content -->
                <div class="flex-1">
                  <div class="flex justify-between items-start">
                    <h4 class="text-lg font-semibold text-gray-800 group-hover:text-[#206bc4] transition-colors">{{ meeting.title }}</h4>
                    <span class="text-sm font-mono text-gray-500 bg-gray-100 px-2 py-1 rounded">{{ meeting.time }}</span>
                  </div>
                  <div class="mt-2 flex items-center text-sm text-gray-500">
                    <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path></svg>
                    Google Meet
                    <span class="mx-2">•</span>
                    <span class="text-gray-500">{{ meeting.attendees.length }} Attendees</span>
                  </div>
                </div>

                <!-- Action -->
                <div class="ml-4">
                  <button class="p-2 hover:bg-gray-100 rounded-full text-gray-400 hover:text-[#206bc4] transition-colors">
                     <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"></path></svg>
                  </button>
                </div>
             </div>
          }
        </div>

        <!-- Mini Calendar & Stats -->
        <div class="space-y-6">
          <div class="bg-white p-6 rounded-xl border border-gray-200 shadow-sm">
            <h3 class="text-gray-800 font-semibold mb-4">Quick Calendar</h3>
            <div class="grid grid-cols-7 gap-2 text-center text-sm mb-2">
              <span class="text-gray-400 text-xs uppercase font-bold">S</span>
              <span class="text-gray-400 text-xs uppercase font-bold">M</span>
              <span class="text-gray-400 text-xs uppercase font-bold">T</span>
              <span class="text-gray-400 text-xs uppercase font-bold">W</span>
              <span class="text-gray-400 text-xs uppercase font-bold">T</span>
              <span class="text-gray-400 text-xs uppercase font-bold">F</span>
              <span class="text-gray-400 text-xs uppercase font-bold">S</span>
            </div>
            <div class="grid grid-cols-7 gap-2 text-center text-sm font-medium">
              <span class="py-1 text-gray-400">29</span>
              <span class="py-1 text-gray-400">30</span>
              <span class="py-1 text-gray-700 hover:bg-gray-100 rounded cursor-pointer">1</span>
              <span class="py-1 text-gray-700 hover:bg-gray-100 rounded cursor-pointer">2</span>
              <span class="py-1 bg-[#206bc4] text-white rounded shadow-sm">3</span>
              <span class="py-1 text-gray-700 hover:bg-gray-100 rounded cursor-pointer">4</span>
              <span class="py-1 text-gray-700 hover:bg-gray-100 rounded cursor-pointer">5</span>
               <!-- More days... simplified for demo -->
               <span class="py-1 text-gray-700">6</span>
               <span class="py-1 text-gray-700">7</span>
               <span class="py-1 text-gray-700 relative">8 <span class="absolute bottom-0 left-1/2 -translate-x-1/2 w-1 h-1 bg-red-500 rounded-full"></span></span>
               <span class="py-1 text-gray-700">9</span>
               <span class="py-1 text-gray-700">10</span>
               <span class="py-1 text-gray-700">11</span>
               <span class="py-1 text-gray-700">12</span>
            </div>
          </div>

          <div class="bg-white p-6 rounded-xl border border-gray-200 shadow-sm relative overflow-hidden">
             <div class="absolute top-0 right-0 p-4 opacity-10">
               <svg class="w-24 h-24 text-[#206bc4]" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-6h2v6zm0-8h-2V7h2v2z"/></svg>
             </div>
             <h3 class="text-gray-800 font-semibold mb-2 relative z-10">Meeting Stats</h3>
             <div class="flex items-center justify-between mb-2 relative z-10">
               <span class="text-gray-500 text-sm">Hours this week</span>
               <span class="text-[#206bc4] font-bold text-xl">12.5h</span>
             </div>
             <div class="w-full bg-gray-100 h-1.5 rounded-full mb-4 relative z-10">
               <div class="bg-purple-500 h-full rounded-full w-[60%]"></div>
             </div>
             <p class="text-xs text-gray-500 relative z-10">You spend 20% less time in meetings compared to last week.</p>
          </div>
        </div>
      </div>
    </div>
  `
})
export class MeetingsComponent {
  dataService = inject(DataService);

  formatDateMonth(iso: string): string {
    return new Date(iso).toLocaleString('default', { month: 'short' });
  }

  formatDateDay(iso: string): string {
    return new Date(iso).getDate().toString();
  }
}