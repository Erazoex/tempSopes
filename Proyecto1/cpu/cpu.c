#include <linux/module.h>
#include <linux/init.h> // Header para los macros module_init y module_exit
#include <linux/sched/signal.h> // for_each_process()
#include <linux/sched.h>
#include <linux/proc_fs.h> // Header necesario porque se usara proc_fs
#include <linux/seq_file.h> // Header para usar la lib seq_file y manejar el archivo en /proc/
#include <linux/mm.h>


MODULE_LICENSE("GPL");
MODULE_AUTHOR("Brian Erazo");
MODULE_DESCRIPTION("Modulo de CPU para proyecto 1, Laboratorio de Sistemas Operativos 1");

struct task_struct *task;           // sched.h para tareas/procesos
struct task_struct *task_child;     // index de tareas secundarias
struct list_head *list;             // lista de cada tarea

static int escribir_archivo(struct seq_file *file_proc, void *v) {
    int running = 0;
    int sleeping = 0;
    int zombie = 0;
    int stopped = 0;
    unsigned long rss;
    unsigned long total_ram_pages;

    total_ram_pages = totalram_pages();
    if (!total_ram_pages) {
        pr_err("No memory available\n");
        return -EINVAL;
    }

    #ifndef CONFIG_MMU
        pr_err("No MMU, cannot calculate RRS\n");
        return -EINVAL;
    #endif

    unsigned long total_cpu_time = jiffies_to_msecs(get_jiffies_64());
    unsigned long total_usage = 0;

    for_each_process(task) {
        unsigned long cpu_time = jiffies_to_msecs(task->utime + task->stime);
        unsigned long cpu_percentage = (cpu_time * 100) / total_cpu_time;
        total_usage += cpu_time;
    }

    seq_printf(file_proc,"{\n\"cpu_total\":%lu,\n", total_cpu_time);
    seq_printf(file_proc, "\"cpu_percentage\":%lu,\n", (total_usage * 100) / total_cpu_time);
    seq_printf(file_proc, "\"processes\":[\n");
    int b = 0;

    for_each_process(task) {
        if (task->mm) {
            rss = get_mm_rss(task->mm) << PAGE_SHIFT;
        } else {
            rss = 0;
        }
        if (b == 0) {
            seq_printf(file_proc, "{");
            b = 1;
        } else {
            seq_printf(file_proc, ",{");
        }
        seq_printf(file_proc, "\"pid\":%d,\n", task->pid);
        seq_printf(file_proc, "\"name\":\"%s\",\n", task->comm);
        seq_printf(file_proc, "\"user\":%u,\n", task->cred->uid);
        seq_printf(file_proc, "\"state\":%u,\n", task->__state);
        int percentage = (rss * 100) / total_ram_pages;
        seq_printf(file_proc, "\"ram\":%d,\n", percentage);
        seq_printf(file_proc, "\"child\":[\n");
        int a = 0;
        list_for_each(list, &(task->children)) {
            task_child = list_entry(list, struct task_struct, sibling);
            if (a != 0) {
                seq_printf(file_proc, ",{");
                seq_printf(file_proc, "\"pid\":%d,\n", task_child->pid);
                seq_printf(file_proc, "\"name\":\"%s\",\n", task_child->comm);
                seq_printf(file_proc, "\"state\":%u,\n", task_child->__state);
                seq_printf(file_proc, "\"pidParent\":%d\n", task->pid);
                seq_printf(file_proc, "}\n");
            } else {
                seq_printf(file_proc, "{");
                seq_printf(file_proc, "\"pid\":%d,\n", task_child->pid);
                seq_printf(file_proc, "\"name\":\"%s\",\n", task_child->comm);
                seq_printf(file_proc, "\"state\":%u,\n", task_child->__state);
                seq_printf(file_proc, "\"pidParent\":%d\n", task->pid);
                seq_printf(file_proc, "}\n");
                a = 1;
            }
        }
        a = 0;
        seq_printf(file_proc, "\n]");
        if (task->__state == 0) {
            running += 1;
        } else if (task->__state == 1) {
            sleeping += 1;
        } else if (task->__state == 4) {
            zombie += 1;
        } else {
            stopped += 1;
        }
        seq_printf(file_proc, "}\n");
    }
    b = 0;
    seq_printf(file_proc, "],\n");
    seq_printf(file_proc, "\"running\":%d,\n", running);
    seq_printf(file_proc, "\"sleeping\":%d,\n", sleeping);
    seq_printf(file_proc, "\"zombie\":%d,\n", zombie);
    seq_printf(file_proc, "\"stopped\":%d,\n", stopped);
    seq_printf(file_proc, "\"total\":%d\n", running + sleeping + zombie + stopped);
    seq_printf(file_proc, "}\n");
    return 0;
}

static int al_abrir(struct inode *inode, struct file *file) {
    return single_open(file, escribir_archivo, NULL);
}

static struct proc_ops operaciones = {
    .proc_open = al_abrir, 
    .proc_read = seq_read
};

static int _insert(void) {
    proc_create("cpu_so1_1s2024", 0, NULL, &operaciones);
    printk(KERN_INFO "201807253\n");
    return 0;
}

static void _remove(void) {
    remove_proc_entry("cpu_so1_1s2024", NULL);
    printk(KERN_INFO "Removing cpu process module\n");
}

module_init(_insert);
module_exit(_remove);
