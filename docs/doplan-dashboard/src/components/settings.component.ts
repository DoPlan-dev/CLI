
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="max-w-4xl mx-auto">
      <h2 class="text-3xl font-bold text-gray-800 mb-8">Settings</h2>

      <!-- Navigation Tabs (Visual only for this demo) -->
      <div class="flex border-b border-gray-200 mb-8 overflow-x-auto">
        <button class="px-6 py-3 text-[#206bc4] border-b-2 border-[#206bc4] font-medium whitespace-nowrap">General</button>
        <button class="px-6 py-3 text-gray-500 hover:text-gray-800 transition-colors whitespace-nowrap">Agents</button>
        <button class="px-6 py-3 text-gray-500 hover:text-gray-800 transition-colors whitespace-nowrap">Notifications</button>
        <button class="px-6 py-3 text-gray-500 hover:text-gray-800 transition-colors whitespace-nowrap">Integrations</button>
      </div>

      <div class="space-y-8">
        
        <!-- Profile Section -->
        <section class="bg-white rounded-xl border border-gray-200 p-6 shadow-sm">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">Profile Information</h3>
          <div class="flex flex-col md:flex-row items-start">
            <img src="images/chatbot.avif" class="w-20 h-20 rounded-full border border-gray-200 mr-6 mb-4 md:mb-0 object-cover">
            <div class="flex-1 grid gap-4 w-full max-w-lg">
              <div class="grid grid-cols-2 gap-4">
                 <div>
                   <label class="block text-xs text-gray-500 mb-1 font-medium">First Name</label>
                   <input type="text" value="Dev" class="w-full bg-white border border-gray-300 rounded px-3 py-2 text-gray-800 text-sm focus:border-[#206bc4] focus:ring-1 focus:ring-[#206bc4] focus:outline-none transition-all">
                 </div>
                 <div>
                   <label class="block text-xs text-gray-500 mb-1 font-medium">Last Name</label>
                   <input type="text" value="Admin" class="w-full bg-white border border-gray-300 rounded px-3 py-2 text-gray-800 text-sm focus:border-[#206bc4] focus:ring-1 focus:ring-[#206bc4] focus:outline-none transition-all">
                 </div>
              </div>
              <div>
                 <label class="block text-xs text-gray-500 mb-1 font-medium">Role</label>
                 <input type="text" value="Java Fullstack Developer" class="w-full bg-white border border-gray-300 rounded px-3 py-2 text-gray-800 text-sm focus:border-[#206bc4] focus:ring-1 focus:ring-[#206bc4] focus:outline-none transition-all">
              </div>
              <div>
                 <label class="block text-xs text-gray-500 mb-1 font-medium">Email</label>
                 <input type="email" value="admin@doplan.io" class="w-full bg-white border border-gray-300 rounded px-3 py-2 text-gray-800 text-sm focus:border-[#206bc4] focus:ring-1 focus:ring-[#206bc4] focus:outline-none transition-all">
              </div>
            </div>
          </div>
        </section>

        <!-- Preferences -->
        <section class="bg-white rounded-xl border border-gray-200 p-6 shadow-sm">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">Preferences</h3>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between py-2 border-b border-gray-100">
              <div>
                <p class="text-gray-800 font-medium">Dark Mode</p>
                <p class="text-xs text-gray-500">Enable system-wide dark theme</p>
              </div>
              <!-- Toggle Switch -->
              <div class="relative inline-flex items-center cursor-pointer">
                 <input type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-[#206bc4]"></div>
              </div>
            </div>

            <div class="flex items-center justify-between py-2 border-b border-gray-100">
              <div>
                <p class="text-gray-800 font-medium">Desktop Notifications</p>
                <p class="text-xs text-gray-500">Get alerted for urgent tasks</p>
              </div>
              <div class="relative inline-flex items-center cursor-pointer">
                 <input type="checkbox" checked class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-[#206bc4]"></div>
              </div>
            </div>

            <div class="flex items-center justify-between py-2">
              <div class="flex items-center gap-3">
                 <img src="images/chatbot.avif" class="w-12 h-12 object-contain hidden sm:block" alt="AI Bot">
                 <div>
                    <p class="text-gray-800 font-medium">AI Suggestions</p>
                    <p class="text-xs text-gray-500">Allow AI agents to suggest task improvements</p>
                 </div>
              </div>
              <div class="relative inline-flex items-center cursor-pointer">
                 <input type="checkbox" class="sr-only peer">
                 <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-[#206bc4]"></div>
              </div>
            </div>
          </div>
        </section>

        <!-- Danger Zone -->
        <section class="border border-red-200 bg-red-50 rounded-xl p-6">
           <h3 class="text-red-600 font-semibold mb-2">Danger Zone</h3>
           <p class="text-sm text-gray-600 mb-4">Once you delete a project, there is no going back. Please be certain.</p>
           <button class="px-4 py-2 bg-white text-red-600 border border-red-200 hover:bg-red-50 hover:border-red-300 rounded text-sm font-medium transition-colors shadow-sm">Delete Project</button>
        </section>

      </div>
    </div>
  `
})
export class SettingsComponent {}
