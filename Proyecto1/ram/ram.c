#include <linux/module.h>
#include <linux/kernel.h> // para usar KERN_INFO
#include <linux/mm.h> // para info_ram
#include <linux/sysinfo.h> // para info del sistema
#include <linux/init.h> // Header para los macros module_init y module_exit
#include <linux/proc_fs.h> // Header necesario porque se usara proc_fs
#include <asm/uaccess.h> // for copy_from_user
#include <linux/seq_file.h> // Header para usar la lib seq_file y manejar el archivo en /proc/


MODULE_LICENSE("GPL");
MODULE_AUTHOR("Brian Erazo");
MODULE_DESCRIPTION("Modulo de RAM para proyecto 1, Laboratorio de Sistemas Operativos 1");

const long megabyte = 1024 * 1024;

// Obteniendo estadisticas del sistema
struct sysinfo si;

static int escribir_archivo(struct seq_file *file_proc, void *v) {
    unsigned long total, used, notused;
    unsigned long proc;
    si_meminfo(&si);

    total = si.totalram * si.mem_unit;
    used = si.freeram * si.mem_unit + si.bufferram * si.mem_unit + si.sharedram * si.mem_unit;
    proc = (used * 100) / total;
    notused = total - used;
    seq_printf(file_proc, "{\n\t\"totalRam\": %lu,\n\t\"usedMemory\": %lu,\n\t\"percentage\": %lu,\n\t\"freeMemory\": %lu\n}", total, used, proc, notused);
    return 0;
}

static int al_abrir(struct inode *inode, struct file *file) {
    return single_open(file, escribir_archivo, NULL);
}

static struct proc_ops operations = {
    .proc_open = al_abrir, 
    .proc_read = seq_read
};

static int _insert(void) {
    proc_create("ram_so1_1s2024", 0, NULL, &operations);
    printk(KERN_INFO "201807253\n");
    return 0;
}

static void _remove(void) {
    remove_proc_entry("ram_so1_1s2024", NULL);
    printk(KERN_INFO "Removing ram process module\n");
}

module_init(_insert);
module_exit(_remove);