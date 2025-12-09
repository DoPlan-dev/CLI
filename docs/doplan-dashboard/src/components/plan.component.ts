import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DataService, Task } from '../services/data.service';

@Component({
  selector: 'app-plan',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="h-full flex flex-col">
      <!-- Toolbar -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-6 gap-4">
        <div>
          <h2 class="text-3xl font-bold text-gray-800">Project Plan</h2>
          <p class="text-gray-500 text-sm mt-1">Manage tasks and track progress across phases.</p>
        </div>
        
        <div class="flex items-center space-x-3 w-full md:w-auto">
          <div class="relative flex-1 md:w-64">
            <input type="text" placeholder="Search tasks..." class="w-full bg-white border border-gray-300 rounded-lg pl-10 pr-4 py-2 text-sm text-gray-700 focus:outline-none focus:border-[#206bc4] focus:ring-1 focus:ring-[#206bc4] transition-colors shadow-sm">
            <svg class="w-4 h-4 text-gray-400 absolute left-3 top-2.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
          </div>
          <button class="p-2 bg-white border border-gray-300 rounded-lg text-gray-500 hover:text-[#206bc4] hover:border-[#206bc4] transition-all shadow-sm">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z"></path></svg>
          </button>
          <button (click)="createNewTask()" class="px-4 py-2 bg-[#206bc4] text-white font-bold rounded-lg text-sm shadow-sm hover:bg-[#1a5c98] transition-all">
            + New Task
          </button>
        </div>
      </div>

      <!-- Kanban Board -->
      <div class="flex-1 overflow-x-auto pb-4">
        <div class="flex gap-6 min-w-[1000px] h-full">
          
          <!-- TODO Column -->
          <div class="flex-1 flex flex-col bg-gray-100/80 rounded-xl border border-gray-200 p-4">
            <div class="flex justify-between items-center mb-4">
              <h3 class="font-semibold text-gray-700 flex items-center">
                <span class="w-2 h-2 rounded-full bg-gray-500 mr-2"></span> To Do
              </h3>
              <span class="bg-white border border-gray-200 text-gray-500 text-xs px-2 py-0.5 rounded-full shadow-sm">{{ getTasksByStatus('todo').length }}</span>
            </div>
            
            <div class="flex-1 space-y-3 overflow-y-auto pr-2 custom-scrollbar">
              @for (task of getTasksByStatus('todo'); track task.id) {
                <div class="bg-white p-4 rounded-lg border border-gray-200 shadow-sm hover:shadow-md hover:border-[#206bc4]/50 transition-all cursor-move group relative">
                  <div class="flex justify-between items-start mb-2">
                    <span class="text-xs font-medium px-2 py-0.5 rounded bg-blue-50 text-blue-600 border border-blue-100">{{ task.tag }}</span>
                    <button class="opacity-0 group-hover:opacity-100 text-gray-400 hover:text-gray-600 transition-opacity">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z"></path></svg>
                    </button>
                  </div>
                  <h4 class="text-gray-800 text-sm font-medium mb-3">{{ task.title }}</h4>
                  <div class="flex justify-between items-center">
                    <div class="flex -space-x-2">
                       <div class="w-6 h-6 rounded-full bg-gray-100 border border-gray-200 flex items-center justify-center text-[10px] text-gray-600 font-bold">{{ task.assignee.charAt(0) }}</div>
                    </div>
                    <div class="flex space-x-2">
                       <button (click)="moveTask(task.id, 'in-progress')" class="p-1 hover:bg-gray-100 rounded text-gray-400 hover:text-[#206bc4]" title="Move to In Progress">
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path></svg>
                       </button>
                    </div>
                  </div>
                </div>
              }
              <button (click)="createNewTask()" class="w-full py-2 border-2 border-dashed border-gray-300 rounded-lg text-sm text-gray-500 hover:text-[#206bc4] hover:border-[#206bc4]/50 hover:bg-white transition-all">
                + Add Card
              </button>
            </div>
          </div>

          <!-- In Progress Column -->
          <div class="flex-1 flex flex-col bg-gray-100/80 rounded-xl border border-gray-200 p-4">
            <div class="flex justify-between items-center mb-4">
              <h3 class="font-semibold text-[#206bc4] flex items-center">
                <span class="w-2 h-2 rounded-full bg-[#206bc4] mr-2"></span> In Progress
              </h3>
              <span class="bg-[#206bc4]/10 border border-[#206bc4]/20 text-[#206bc4] text-xs px-2 py-0.5 rounded-full shadow-sm">{{ getTasksByStatus('in-progress').length }}</span>
            </div>
            
             <div class="flex-1 space-y-3 overflow-y-auto pr-2 custom-scrollbar">
              @for (task of getTasksByStatus('in-progress'); track task.id) {
                <div class="bg-white p-4 rounded-lg border border-l-4 border-l-[#206bc4] border-gray-200 shadow-sm hover:shadow-md transition-all cursor-move group">
                   <div class="flex justify-between items-start mb-2">
                    <span class="text-xs font-medium px-2 py-0.5 rounded bg-purple-50 text-purple-600 border border-purple-100">{{ task.tag }}</span>
                  </div>
                  <h4 class="text-gray-800 text-sm font-medium mb-3">{{ task.title }}</h4>
                  <div class="w-full bg-gray-100 h-1 mb-3 rounded-full overflow-hidden">
                     <div class="bg-[#206bc4] h-full w-2/3"></div>
                  </div>
                  <div class="flex justify-between items-center">
                    <div class="text-xs text-gray-500">Due Tomorrow</div>
                    <button (click)="moveTask(task.id, 'done')" class="p-1 hover:bg-gray-100 rounded text-gray-400 hover:text-green-600" title="Move to Done">
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
                    </button>
                  </div>
                </div>
              }
            </div>
          </div>

          <!-- Done Column -->
          <div class="flex-1 flex flex-col bg-gray-100/80 rounded-xl border border-gray-200 p-4">
            <div class="flex justify-between items-center mb-4">
              <h3 class="font-semibold text-green-600 flex items-center">
                <span class="w-2 h-2 rounded-full bg-green-500 mr-2"></span> Done
              </h3>
              <span class="bg-green-100 border border-green-200 text-green-600 text-xs px-2 py-0.5 rounded-full shadow-sm">{{ getTasksByStatus('done').length }}</span>
            </div>
            
             <div class="flex-1 space-y-3 overflow-y-auto pr-2 custom-scrollbar">
              @for (task of getTasksByStatus('done'); track task.id) {
                <div class="bg-white p-4 rounded-lg border border-gray-200 opacity-70 hover:opacity-100 transition-all shadow-sm">
                   <div class="flex justify-between items-start mb-2">
                    <span class="text-xs font-medium px-2 py-0.5 rounded bg-green-50 text-green-600 border border-green-100">Completed</span>
                  </div>
                  <h4 class="text-gray-500 text-sm font-medium mb-3 line-through">{{ task.title }}</h4>
                  <div class="flex justify-end">
                     <svg class="w-5 h-5 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                  </div>
                </div>
              }
            </div>
          </div>

        </div>
      </div>
    </div>
  `
})
export class PlanComponent {
  dataService = inject(DataService);

  getTasksByStatus(status: string): Task[] {
    return this.dataService.tasks().filter(t => t.status === status);
  }

  moveTask(id: string, status: 'todo' | 'in-progress' | 'done') {
    this.dataService.moveTask(id, status);
  }

  createNewTask() {
    this.dataService.addTask('New Untitled Task', 'General');
  }
}