import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavbarComponent } from './components/sidebar.component'; // Import from sidebar file which now contains Navbar
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, NavbarComponent, CommonModule],
  template: `
    <div class="min-h-screen bg-[#f4f6fa] text-[#1e293b] flex flex-col font-sans">
      <app-navbar></app-navbar>
      
      <main class="flex-1 w-full max-w-[1600px] mx-auto px-4 sm:px-6 lg:px-8 py-8 animate-in fade-in duration-500">
        <router-outlet></router-outlet>
      </main>

      <footer class="border-t border-gray-200 mt-auto bg-white py-8">
        <div class="max-w-[1600px] mx-auto px-4 text-center text-sm text-gray-500">
          <p>&copy; 2024 DoPlan Dashboard. Inspired by <a href="https://tabler.io" class="text-[#206bc4] hover:underline" target="_blank">Tabler.io</a>.</p>
        </div>
      </footer>
    </div>
  `
})
export class AppComponent {}