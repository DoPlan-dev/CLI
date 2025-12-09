
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink, RouterLinkActive } from '@angular/router';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [CommonModule, RouterLink, RouterLinkActive],
  template: `
    <header class="bg-white border-b border-gray-200 sticky top-0 z-50 shadow-sm">
      <div class="max-w-[1600px] mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex h-16 items-center justify-between">
          
          <!-- Left: Logo & Nav -->
          <div class="flex items-center gap-8">
            <!-- Logo -->
            <a routerLink="/" class="flex items-center gap-2 hover:opacity-80 transition-opacity">
              <span class="text-xl font-extrabold tracking-tight text-[#206bc4]">DoPlan.dev</span>
            </a>

            <!-- Desktop Nav -->
            <nav class="hidden md:flex items-center gap-1">
              <a routerLink="/" 
                 routerLinkActive="text-[#206bc4] bg-[#206bc4]/10 font-semibold" 
                 [routerLinkActiveOptions]="{exact: true}"
                 class="px-3 py-2 rounded-md text-sm font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-100 transition-all flex items-center gap-2">
                 <svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12l-2 0l9 -9l9 9l-2 0" /><path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7" /><path d="M9 21v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6" /></svg>
                 Dashboard
              </a>
              
              <a routerLink="/plan" 
                 routerLinkActive="text-[#206bc4] bg-[#206bc4]/10 font-semibold"
                 class="px-3 py-2 rounded-md text-sm font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-100 transition-all flex items-center gap-2">
                 <svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 3l0 18" /><path d="M6 9l12 0" /><path d="M6 15l12 0" /></svg>
                 Plan
              </a>

              <a routerLink="/meetings" 
                 routerLinkActive="text-[#206bc4] bg-[#206bc4]/10 font-semibold"
                 class="px-3 py-2 rounded-md text-sm font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-100 transition-all flex items-center gap-2">
                 <svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 5m0 2a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2z" /><path d="M16 3l0 4" /><path d="M8 3l0 4" /><path d="M4 11l16 0" /><path d="M8 15h2v2h-2z" /></svg>
                 Meetings
              </a>

              <a routerLink="/achievements" 
                 routerLinkActive="text-[#206bc4] bg-[#206bc4]/10 font-semibold"
                 class="px-3 py-2 rounded-md text-sm font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-100 transition-all flex items-center gap-2">
                 <svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 5l0 14" /><path d="M18 13l-6 6" /><path d="M6 13l6 6" /></svg>
                 Achievements
              </a>

               <a routerLink="/settings" 
                 routerLinkActive="text-[#206bc4] bg-[#206bc4]/10 font-semibold"
                 class="px-3 py-2 rounded-md text-sm font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-100 transition-all flex items-center gap-2">
                 <svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10.325 4.317c.426 -1.756 2.924 -1.756 3.35 0a1.724 1.724 0 0 0 2.573 1.066c1.543 -.94 3.31 .826 2.37 2.37a1.724 1.724 0 0 0 1.065 2.572c1.756 .426 1.756 2.924 0 3.35a1.724 1.724 0 0 0 -1.066 2.573c.94 1.543 -.826 3.31 -2.37 2.37a1.724 1.724 0 0 0 -2.572 1.065c-.426 1.756 -2.924 1.756 -3.35 0a1.724 1.724 0 0 0 -2.573 -1.066c-1.543 .94 -3.31 -.826 -2.37 -2.37a1.724 1.724 0 0 0 -1.065 -2.572c-1.756 -.426 -1.756 -2.924 0 -3.35a1.724 1.724 0 0 0 1.066 -2.573c-.94 -1.543 .826 -3.31 2.37 -2.37c.996 .608 2.296 .07 2.572 -1.065z" /><path d="M9 12a3 3 0 1 0 6 0a3 3 0 0 0 -6 0" /></svg>
                 Settings
              </a>
            </nav>
          </div>

          <!-- Right: Search & Profile -->
          <div class="flex items-center gap-4">
             <!-- Search (Desktop) -->
            <div class="hidden lg:block relative">
              <input type="text" placeholder="Search..." class="bg-gray-100 border border-gray-200 rounded-full py-1.5 pl-9 pr-4 text-sm text-gray-700 focus:outline-none focus:border-[#206bc4] focus:ring-1 focus:ring-[#206bc4] w-64 transition-all placeholder:text-gray-500">
              <svg class="w-4 h-4 text-gray-500 absolute left-3 top-2.5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10 10m-7 0a7 7 0 1 0 14 0a7 7 0 1 0 -14 0" /><path d="M21 21l-6 -6" /></svg>
            </div>

            <!-- Notifications -->
            <button class="relative p-2 text-gray-500 hover:text-[#206bc4] transition-colors">
              <svg class="w-6 h-6" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6" /><path d="M9 17v1a3 3 0 0 0 6 0v-1" /></svg>
              <span class="absolute top-2 right-2 w-2 h-2 bg-red-500 rounded-full border-2 border-white"></span>
            </button>

            <!-- Profile -->
            <div class="flex items-center gap-3 pl-4 border-l border-gray-200">
              <div class="text-right hidden sm:block">
                <p class="text-sm font-medium text-gray-700 leading-none">Dev Admin</p>
                <p class="text-xs text-gray-500 mt-1 leading-none">Project Lead</p>
              </div>
              <div class="h-9 w-9 rounded-full bg-gray-100 p-[1px] cursor-pointer hover:shadow-md transition-all overflow-hidden">
                <img src="images/chatbot.avif" alt="User" class="rounded-full h-full w-full object-cover border border-gray-200">
              </div>
            </div>

             <!-- Mobile Menu Button -->
            <button class="md:hidden p-2 text-gray-500 hover:text-[#206bc4]" (click)="toggleMobileMenu()">
               <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Mobile Menu (Expandable) -->
      @if (isMobileMenuOpen) {
        <div class="md:hidden border-t border-gray-200 bg-white animate-in slide-in-from-top-2 duration-200 shadow-lg">
          <div class="px-2 pt-2 pb-3 space-y-1">
             <a routerLink="/" (click)="closeMobileMenu()" class="block px-3 py-2 rounded-md text-base font-medium text-gray-600 hover:text-[#206bc4] hover:bg-gray-50">Dashboard</a>
             <a routerLink="/plan" (click)="closeMobileMenu()" class="block px-3 py-2 rounded-md text-base font-medium text-gray-600 hover:text-[#206bc4] hover:bg-gray-50">Plan</a>
             <a routerLink="/meetings" (click)="closeMobileMenu()" class="block px-3 py-2 rounded-md text-base font-medium text-gray-600 hover:text-[#206bc4] hover:bg-gray-50">Meetings</a>
             <a routerLink="/achievements" (click)="closeMobileMenu()" class="block px-3 py-2 rounded-md text-base font-medium text-gray-600 hover:text-[#206bc4] hover:bg-gray-50">Achievements</a>
             <a routerLink="/settings" (click)="closeMobileMenu()" class="block px-3 py-2 rounded-md text-base font-medium text-gray-600 hover:text-[#206bc4] hover:bg-gray-50">Settings</a>
          </div>
        </div>
      }
    </header>
  `
})
export class NavbarComponent {
  isMobileMenuOpen = false;

  toggleMobileMenu() {
    this.isMobileMenuOpen = !this.isMobileMenuOpen;
  }

  closeMobileMenu() {
    this.isMobileMenuOpen = false;
  }
}
