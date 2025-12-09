import { Injectable, signal, computed } from '@angular/core';

export interface Task {
  id: string;
  title: string;
  status: 'todo' | 'in-progress' | 'done';
  tag: string;
  assignee: string;
}

export interface Meeting {
  id: string;
  title: string;
  date: string; // ISO string
  time: string;
  attendees: string[];
}

export interface Achievement {
  id: string;
  title: string;
  description: string;
  icon: string;
  unlocked: boolean;
  progress: number; // 0-100
}

@Injectable({
  providedIn: 'root'
})
export class DataService {
  // State Signals
  tasks = signal<Task[]>([
    { id: '1', title: 'Migrate to Java 21', status: 'in-progress', tag: 'Backend', assignee: 'Alex' },
    { id: '2', title: 'Refactor HTML Structure', status: 'todo', tag: 'Frontend', assignee: 'Sam' },
    { id: '3', title: 'Spring Boot Init', status: 'done', tag: 'Backend', assignee: 'Jordan' },
    { id: '4', title: 'Fix CSS Grid Layout', status: 'in-progress', tag: 'Frontend', assignee: 'Alex' },
    { id: '5', title: 'Hibernate Query Optimization', status: 'todo', tag: 'Database', assignee: 'Sam' },
    { id: '6', title: 'Implement Semantic HTML', status: 'todo', tag: 'Accessibility', assignee: 'Casey' },
  ]);

  meetings = signal<Meeting[]>([
    { id: '1', title: 'JVM Tuning Sync', date: new Date().toISOString(), time: '09:00 AM', attendees: ['Alex', 'Sam', 'Jordan'] },
    { id: '2', title: 'Frontend Architecture', date: new Date(Date.now() + 86400000).toISOString(), time: '02:00 PM', attendees: ['Team'] },
    { id: '3', title: 'Client Demo (JSP vs HTML)', date: new Date(Date.now() - 86400000).toISOString(), time: '11:00 AM', attendees: ['Alex', 'Client'] },
  ]);

  achievements = signal<Achievement[]>([
    { id: '1', title: 'Syntax Sorcerer', description: 'Write error-free Java code for a day', icon: 'sun', unlocked: true, progress: 100 },
    { id: '2', title: 'Garbage Collector', description: 'Refactor 10 legacy classes', icon: 'bug', unlocked: true, progress: 100 },
    { id: '3', title: 'Pixel Perfect', description: 'Match HTML to Figma exactly', icon: 'flame', unlocked: false, progress: 60 },
    { id: '4', title: 'Code Reviewer', description: 'Review 50 Pull Requests', icon: 'users', unlocked: false, progress: 45 },
  ]);

  // Computed Values
  totalTasks = computed(() => this.tasks().length);
  completedTasks = computed(() => this.tasks().filter(t => t.status === 'done').length);
  pendingTasks = computed(() => this.tasks().filter(t => t.status === 'todo').length);
  inProgressTasks = computed(() => this.tasks().filter(t => t.status === 'in-progress').length);
  
  progressPercentage = computed(() => {
    const total = this.totalTasks();
    if (total === 0) return 0;
    return Math.round((this.completedTasks() / total) * 100);
  });

  // Actions
  addTask(title: string, tag: string) {
    this.tasks.update(tasks => [
      ...tasks, 
      { 
        id: Math.random().toString(36).substr(2, 9), 
        title, 
        status: 'todo', 
        tag, 
        assignee: 'Me' 
      }
    ]);
  }

  moveTask(id: string, status: 'todo' | 'in-progress' | 'done') {
    this.tasks.update(tasks => 
      tasks.map(t => t.id === id ? { ...t, status } : t)
    );
  }

  toggleAchievement(id: string) {
    this.achievements.update(achievements =>
      achievements.map(a => a.id === id ? { ...a, unlocked: !a.unlocked } : a)
    );
  }
}