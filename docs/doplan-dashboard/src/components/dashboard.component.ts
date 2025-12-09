
import { Component, inject, ViewChild, ElementRef, AfterViewInit, OnDestroy } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DataService } from '../services/data.service';

declare var Chart: any;

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="space-y-8">
      
      <!-- Welcome Hero with Illustration -->
      <div class="bg-white rounded-xl border border-gray-200 p-8 relative overflow-hidden flex flex-col md:flex-row items-center justify-between shadow-sm">
         <div class="relative z-10 max-w-2xl">
            <h2 class="text-3xl font-bold text-gray-800 mb-2">Welcome back, Java Dev!</h2>
            <p class="text-gray-500 mb-6">You have completed <span class="text-[#206bc4] font-bold">{{ dataService.completedTasks() }} tasks</span> in the migration project. The HTML refactor is on schedule.</p>
            <div class="flex gap-3">
               <button class="px-5 py-2.5 bg-[#206bc4] hover:bg-[#1a5c98] text-white font-bold rounded-lg shadow-sm hover:shadow transition-all">
                  Go to Plan
               </button>
               <button class="px-5 py-2.5 bg-white hover:bg-gray-50 text-gray-700 font-medium rounded-lg border border-gray-300 transition-all">
                  View Reports
               </button>
            </div>
         </div>
         <!-- Decorative Illustration -->
         <div class="hidden md:block relative w-64 h-48">
            <img src="images/chatbot.avif" class="w-full h-full object-contain" alt="Dashboard Illustration">
         </div>
      </div>

      <!-- Stats Grid -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
         
         <!-- Progress Card -->
         <div class="bg-white p-6 rounded-xl border border-gray-200 shadow-sm hover:border-[#206bc4]/50 transition-all group">
            <div class="flex items-center justify-between mb-4">
              <span class="text-gray-500 text-sm font-medium">Sprint Progress</span>
              <span class="bg-green-100 text-green-600 text-xs px-2 py-1 rounded flex items-center gap-1">
                 <svg class="w-3 h-3" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="23 6 13.5 15.5 8.5 10.5 1 18"></polyline><polyline points="17 6 23 6 23 12"></polyline></svg>
                 +12%
              </span>
            </div>
            <div class="text-3xl font-bold text-gray-800 mb-2">{{ dataService.progressPercentage() }}%</div>
            <div class="w-full bg-gray-100 h-1.5 rounded-full overflow-hidden">
               <div class="bg-[#206bc4] h-full rounded-full" [style.width.%]="dataService.progressPercentage()"></div>
            </div>
         </div>

         <!-- Tasks Card -->
         <div class="bg-white p-6 rounded-xl border border-gray-200 shadow-sm hover:border-[#206bc4]/50 transition-all group">
            <div class="flex items-center justify-between mb-4">
              <span class="text-gray-500 text-sm font-medium">Pending Tasks</span>
              <span class="bg-yellow-100 text-yellow-600 text-xs px-2 py-1 rounded">-2%</span>
            </div>
            <div class="text-3xl font-bold text-gray-800 mb-1">{{ dataService.pendingTasks() }}</div>
            <div class="text-xs text-gray-500 flex items-center gap-1">
               <span class="w-2 h-2 rounded-full bg-yellow-500"></span> Needs attention
            </div>
         </div>

         <!-- Deadlines Card -->
         <div class="bg-white p-6 rounded-xl border border-gray-200 shadow-sm hover:border-[#206bc4]/50 transition-all group">
            <div class="flex items-center justify-between mb-4">
               <span class="text-gray-500 text-sm font-medium">Release Date</span>
            </div>
            <div class="text-3xl font-bold text-gray-800 mb-1">4 Days</div>
            <div class="text-xs text-gray-500">Java 21 Rollout</div>
         </div>

         <!-- Velocity Card -->
         <div class="bg-white p-6 rounded-xl border border-gray-200 shadow-sm hover:border-[#206bc4]/50 transition-all group">
            <div class="flex items-center justify-between mb-2">
               <span class="text-gray-500 text-sm font-medium">Commit Velocity</span>
            </div>
            <div class="flex items-end gap-1 h-12">
               <div class="bg-gray-100 w-full rounded-t-sm h-[40%] group-hover:bg-[#206bc4]/30 transition-colors"></div>
               <div class="bg-gray-100 w-full rounded-t-sm h-[60%] group-hover:bg-[#206bc4]/40 transition-colors"></div>
               <div class="bg-gray-100 w-full rounded-t-sm h-[30%] group-hover:bg-[#206bc4]/30 transition-colors"></div>
               <div class="bg-gray-100 w-full rounded-t-sm h-[80%] group-hover:bg-[#206bc4]/50 transition-colors"></div>
               <div class="bg-gray-100 w-full rounded-t-sm h-[50%] group-hover:bg-[#206bc4]/40 transition-colors"></div>
               <div class="bg-[#206bc4] w-full rounded-t-sm h-[90%]"></div>
               <div class="bg-gray-100 w-full rounded-t-sm h-[70%] group-hover:bg-[#206bc4]/60 transition-colors"></div>
            </div>
         </div>
      </div>

      <!-- Content Grid -->
      <div class="grid grid-cols-1 xl:grid-cols-3 gap-6">
         <!-- Chart -->
         <div class="xl:col-span-2 bg-white rounded-xl border border-gray-200 p-6 min-h-[400px] shadow-sm flex flex-col">
            <div class="flex justify-between items-center mb-6">
               <h3 class="text-lg font-semibold text-gray-800">Code Frequency</h3>
               <button class="text-gray-400 hover:text-gray-600 p-1 hover:bg-gray-100 rounded">
                  <svg class="w-5 h-5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0 -18 0" /><path d="M12 9h.01" /><path d="M11 12h1v4h1" /></svg>
               </button>
            </div>
            <!-- Canvas Chart -->
            <div class="relative w-full flex-1 min-h-[300px]">
                 <canvas #chartCanvas></canvas>
            </div>
         </div>

         <!-- Side List -->
         <div class="bg-white rounded-xl border border-gray-200 overflow-hidden flex flex-col shadow-sm">
            <div class="p-6 border-b border-gray-100">
               <h3 class="text-lg font-semibold text-gray-800">Recent Commits</h3>
            </div>
            <div class="flex-1 overflow-y-auto max-h-[400px]">
                <div class="p-4 flex gap-4 hover:bg-gray-50 transition-colors border-b border-gray-100">
                   <div class="w-2 h-2 mt-2 rounded-full bg-blue-500 shrink-0"></div>
                   <div>
                      <p class="text-sm text-gray-700"><strong>Alex</strong> pushed to <code>feature/java-21</code></p>
                      <p class="text-xs text-gray-500 mt-0.5">2 mins ago</p>
                   </div>
                </div>
                <div class="p-4 flex gap-4 hover:bg-gray-50 transition-colors border-b border-gray-100">
                   <div class="w-2 h-2 mt-2 rounded-full bg-green-500 shrink-0"></div>
                   <div>
                      <p class="text-sm text-gray-700"><strong>System</strong> deployed to staging</p>
                      <p class="text-xs text-gray-500 mt-0.5">1 hour ago</p>
                   </div>
                </div>
                <div class="p-4 flex gap-4 hover:bg-gray-50 transition-colors border-b border-gray-100">
                   <div class="w-2 h-2 mt-2 rounded-full bg-purple-500 shrink-0"></div>
                   <div>
                      <p class="text-sm text-gray-700"><strong>Sarah</strong> updated HTML templates</p>
                      <p class="text-xs text-gray-500 mt-0.5">3 hours ago</p>
                   </div>
                </div>
                 <div class="p-4 flex gap-4 hover:bg-gray-50 transition-colors border-b border-gray-100">
                   <div class="w-2 h-2 mt-2 rounded-full bg-yellow-500 shrink-0"></div>
                   <div>
                      <p class="text-sm text-gray-700"><strong>Jira</strong> synced 3 bug reports</p>
                      <p class="text-xs text-gray-500 mt-0.5">Yesterday</p>
                   </div>
                </div>
            </div>
            <div class="p-4 bg-gray-50 text-center border-t border-gray-200">
               <a href="#" class="text-sm text-[#206bc4] hover:underline font-medium">View all activity</a>
            </div>
         </div>
      </div>

    </div>
  `
})
export class DashboardComponent implements AfterViewInit, OnDestroy {
  dataService = inject(DataService);
  
  @ViewChild('chartCanvas') chartCanvas!: ElementRef<HTMLCanvasElement>;
  chartInstance: any;

  ngAfterViewInit() {
    this.initChart();
  }

  ngOnDestroy() {
    if (this.chartInstance) {
      this.chartInstance.destroy();
    }
  }

  initChart() {
    if (typeof Chart === 'undefined') {
      console.warn('Chart.js not loaded');
      return;
    }

    const ctx = this.chartCanvas.nativeElement.getContext('2d');
    if (!ctx) return;

    // Create Gradient for the fill
    const gradient = ctx.createLinearGradient(0, 0, 0, 300);
    gradient.addColorStop(0, 'rgba(32, 107, 196, 0.15)'); // Tabler Blue transparent
    gradient.addColorStop(1, 'rgba(32, 107, 196, 0.0)');

    this.chartInstance = new Chart(ctx, {
      type: 'line',
      data: {
        labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        datasets: [{
          label: 'Lines of Code',
          data: [150, 230, 180, 320, 250, 140, 350],
          borderColor: '#206bc4',
          backgroundColor: gradient,
          borderWidth: 2,
          pointBackgroundColor: '#ffffff',
          pointBorderColor: '#206bc4',
          pointBorderWidth: 2,
          pointRadius: 4,
          pointHoverRadius: 6,
          fill: true,
          tension: 0.4 // Bezier curve
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false
          },
          tooltip: {
            backgroundColor: '#1e293b',
            titleColor: '#ffffff',
            bodyColor: '#cbd5e1',
            borderColor: '#e2e8f0',
            borderWidth: 1,
            padding: 12,
            displayColors: false,
            callbacks: {
              label: function(context: any) {
                return context.parsed.y + ' LoC';
              }
            }
          }
        },
        scales: {
          y: {
            beginAtZero: true,
            grid: {
              color: '#f1f5f9',
              drawBorder: false,
            },
            ticks: {
              color: '#64748b',
              font: {
                family: 'Inter',
                size: 11
              }
            }
          },
          x: {
            grid: {
              display: false,
              drawBorder: false,
            },
            ticks: {
              color: '#64748b',
              font: {
                family: 'Inter',
                size: 11
              }
            }
          }
        },
        interaction: {
          intersect: false,
          mode: 'index',
        },
        elements: {
          point: {
            radius: 0,
            hitRadius: 10,
            hoverRadius: 4
          }
        }
      }
    });
  }
}
